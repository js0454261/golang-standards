package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"

	"os"
)

var (

	//stringconexao é a string que conecta ao banco de dados
	StringConexadoBanco = ""

	//nesta porta vai ser onde a api vai tá carregando
	Porta = 0
)

// carregar vai inicializar as variaveis de ambientes
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {

		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}

	StringConexadoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=local",

		os.Getenv("DB_USUARIOS"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
