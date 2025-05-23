package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/admin"
	"github.com/labstack/echo/v4"
	"github.com/stackus/hxgo/hxecho"
)

func (a *App) adminRoutes(e *echo.Group) *echo.Group {
	adminGroup := e.Group("/admin")
	adminGroup.Use(a.ValidateAdminMiddleware)

	adminGroup.GET("", a.adminRootHandler).Name = "admin"
	adminGroup.POST("/config", a.adminConfigUpdateHandler).Name = "admin-config-update"

	adminUsersGroup := adminGroup.Group("/users")
	adminUsersGroup.GET("/:id/edit", a.adminUserEditHandler).Name = "admin-user-edit"
	adminUsersGroup.POST("/:id", a.adminUserUpdateHandler).Name = "admin-user-update"
	adminUsersGroup.GET("/:id/delete", a.adminUserDeleteConfirmHandler).Name = "admin-user-delete-confirm"
	adminUsersGroup.POST("/:id/delete", a.adminUserDeleteHandler).Name = "admin-user-delete"
	adminUsersGroup.GET("/:id", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, a.echo.Reverse("admin-user-edit", c.Param("id")))
	}).Name = "admin-user-show"

	return adminGroup
}

func (a *App) adminRootHandler(c echo.Context) error {
	users, err := database.GetUsers(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, admin.Root(users))
}

func (a *App) adminUserEditHandler(c echo.Context) error {
	user, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, "/admin", err)
	}

	if user == nil {
		return a.redirectWithError(c, "/admin", ErrUserNotFound)
	}

	return Render(c, http.StatusOK, admin.EditUser(user))
}

func (a *App) adminUserUpdateHandler(c echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, "/admin", err)
	}

	u.Name = c.FormValue("name")
	u.Username = c.FormValue("username")
	u.Admin = isChecked(c.FormValue("admin"))
	u.Active = isChecked(c.FormValue("active"))

	if c.FormValue("password") != "" {
		if err := u.SetPassword(c.FormValue("password")); err != nil {
			return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
		}
	}

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_user_s_has_been_updated", u.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin-user-show", c.Param("id")))
}

func (a *App) adminUserDeleteHandler(c echo.Context) error { //nolint:dupl
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := u.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin-user-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_user_s_has_been_deleted", u.Name)

	if hxecho.IsHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.echo.Reverse("admin"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin"))
}

func (a *App) adminConfigUpdateHandler(c echo.Context) error {
	var cnf database.Config

	if err := c.Bind(&cnf); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := cnf.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	if err := a.ResetConfiguration(); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	a.addNoticeT(c, "Config updated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("admin"))
}

func isChecked(value string) bool {
	return value == "on"
}

func (a *App) adminUserDeleteConfirmHandler(c echo.Context) error {
	u, err := a.getUser(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("admin"), err)
	}

	return Render(c, http.StatusOK, admin.DeleteModal(u))
}
