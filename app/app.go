package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar - vai retornar a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome do servidor na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
		{
			Name:   "address",
			Usage:  "Busca o endereço do host na internet",
			Flags:  flags,
			Action: buscarEndereco,
		},
		{
			Name:   "cname",
			Usage:  "Busca o nome canonico do host na internet",
			Flags:  flags,
			Action: buscarNome,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func buscarEndereco(c *cli.Context) {
	host := c.String("host")

	enderecos, erro := net.LookupAddr(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, endereco := range enderecos {
		fmt.Println(endereco)
	}
}

func buscarNome(c *cli.Context) {
	host := c.String("host")

	names, erro := net.LookupCNAME(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
