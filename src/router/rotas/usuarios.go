package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{
	{

		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},

	{

		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},

	{

		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuariosId,
		RequerAutenticacao: false,
	},

	{

		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUsuarios,
		RequerAutenticacao: false,
	},

	{

		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DeletarUsuarios,
		RequerAutenticacao: false,
	},
}
