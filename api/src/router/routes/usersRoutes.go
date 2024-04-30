package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/usuarios",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUserById,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
