package main

import (
	"Golanta/routeur"
	"Golanta/templates"
)

func main() {
	templates.InitTemplate()
	routeur.Initserv()
}
