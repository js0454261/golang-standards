package controllers

import "net/http"

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criar usuarios"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuarios"))
}

func BuscarUsuariosId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuarios por id"))
}

func EditarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editar usuarios"))
}

func DeletarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletrar usuarios"))
}
