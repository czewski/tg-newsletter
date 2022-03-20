package constants

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadKey(which string) (token string) {
	Keys := Keys{}

	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("Deu ruim")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	err = json.Unmarshal(byteValue, &Keys)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	if which == "botKey" {
		token = Keys.BotKey
	} else if which == "newsKey" {
		token = Keys.NewsKey
	} else {
		fmt.Println("Deu Ruim")
	}

	return token
}

type Keys struct {
	BotKey  string `json:"botKey"`
	NewsKey string `json:"newsKey"`
}

type NewsResult struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      interface{} `json:"author"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	URLToImage  string      `json:"urlToImage"`
	PublishedAt string      `json:"publishedAt"`
	Content     string      `json:"content"`
}

type MessageToSend struct {
	ChatID                int    `json:"chat_id"`
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	ParseMode             string `json:"parse_mode"`
}

const Newsss = `{
	"status": "ok",
	"totalresults": 6,
	"articles": [
	{
	"source": {
	"id": "globo",
	"name": "globo"
	},
	"author": null,
	"title": "bb (bbas3) busca fechar distancia para concorrencia",
	"description": "controlado pelo governo federal, o banco do brasil (bbas3) tem recomendacao de compra, com preco-alvo de r$ 50 para o safra, banco do brasil apresenta forte crescimento de credito, boa qualidade da carteira e esta pronto para apresentar numeros fortes para 20... ",
	"url": "https: //valor. Globo. Com/patrocinado/investe-safra/noticia/2022/03/15/bb-bbas3-busca-fechar-distancia-para-concorrencia. Ghtml",
	"urltoimage": "https: //s2. Glbimg. Com/yxrdbxyefytssdlac8pssbharei=/1200x/smart/filters: cover(): strip_icc()/i. S3. Glbimg. Com/v1/auth_63b422c2caee4269b8b34177e8876b93/internal_photos/bs/2022/e/m/wltwnwtdatzc4equv7qg/gettyimages-1314788516. Jpg",
	"publishedat": "2022-03-15t16:19:47z",
	"content": "em apresentacao virtual, o banco do brasil (bbas3) reiterou a confianca para cumprir o guidance para 2022 (traduzindo para um lucro liquido entre r$23 e r$26 bilhoes), impulsionado principalmente pel... [+3753 chars]"
	},
	{
	"source": {
	"id": null,
	"name": "ig. Com. Br"
	},
	"author": "economia@igcorp. Com. Br (1bilhao)",
	"title": "tres fatores que vao fazer essa acao subir forte nessa semana",
	"description": "acao do banco do brasil fez um lindo candle, com volume e se superar um patamar determinado, vai voar",
	"url": "https: //economia. Ig. Com. Br/1bilhao/2022-03-10/tres-fatores-que-vao-fazer-essa-acao-subir-forte-nessa-semana. Html",
	"urltoimage": "https: //i0. Statig. Com. Br/bancodeimagens/aj/jp/0n/ajjp0n9ypoxj52jp4yn1jb8ek. Jpg",
	"publishedat": "2022-03-10t23:22:48z",
	"content": "bbas3\r\nfonte plataforma profitpro \r\no setor financeiro, assim como a bolsa teve uma boa recuperacao na data de 09/03/2022. \r\ndepois de toda queda por conta do conflito do leste europeu e das sancoes ... [+2339 chars]"
	},
	{
	"source": {
	"id": null,
	"name": "ig. Com. Br"
	},
	"author": "economia@igcorp. Com. Br (1bilhao)",
	"title": "acao de grande empresa brasileira pode disparar mais de 40%, diz safra",
	"description": "acoes do banco do brasil estao operando com um grande desconto e podem ter alta significativa nos proximos meses",
	"url": "https: //economia. Ig. Com. Br/1bilhao/2022-03-09/acao-de-grande-empresa-brasileira-pode-disparar-mais-de-40---diz-safra. Html",
	"urltoimage": "https: //i0. Statig. Com. Br/bancodeimagens/8v/wd/0i/8vwd0itjc909qvm9z4vgc0fk4. Jpg",
	"publishedat": "2022-03-09t23:15:27z",
	"content": "o banco do brasil (bbas3) esta muito barato na b3, alerta o safra em relatorio. Os papeis ja tiveram alta de 20% neste ano, mas o banco ainda acredita em uma alta ate os r$ 50, o que representa ganho... [+2083 chars]"
	},
	{
	"source": {
	"id": null,
	"name": "ig. Com. Br"
	},
	"author": "economia@igcorp. Com. Br (1bilhao)",
	"title": "nubank esta prestes a ser o 4o mais valioso do brasil com queda de 10%",
	"description": "itau unibanco e bradesco ja estao na frente do nubank, enquanto santander brasil esta \"quase la\"",
	"url": "https: //economia. Ig. Com. Br/1bilhao/2022-03-14/nubank-esta-prestes-a-ser-o-4-mais-valioso-do-brasil-com-queda-de-10-. Html",
	"urltoimage": "https: //i0. Statig. Com. Br/bancodeimagens/9h/78/ly/9h78ly8jxfpxxp1yydhpcwtcu. Jpg",
	"publishedat": "2022-03-14t23:51:27z",
	"content": "o nubank (nubr33) abriu o capital em dezembro como o banco brasileiro mais valioso de todos. Menos de quatro meses depois, o nubank esta bem proximo de se tornar o quarto banco com maior capitalizaca... [+1820 chars]"
	},
	{
	"source": {
	"id": null,
	"name": "capitalist. Com. Br"
	},
	"author": "capitalist",
	"title": "bancos agitam o mercado financeiro",
	"description": "alguns bancos destacaram-se na segunda feira (7) no mercado devido a distintas movimentacoes. Confira aqui! ",
	"url": "https: //capitalist. Com. Br/bancos-agitam-o-mercado-financeiro/",
	"urltoimage": "https: //capitalist. Com. Br/wp-content/uploads/2022/03/investimentos-1000x600. Jpg",
	"publishedat": "2022-03-10t20:06:07z",
	"content": "o mercado financeiro, de forma simplificada, e todo o universo de operacoes de compra e venda de ativos desde titulos emitidos por bancos ate acoes de empresas, passando pelos milhares de tipos de ap... [+2768 chars]"
	},
	{
	"source": {
	"id": "globo",
	"name": "globo"
	},
	"author": null,
	"title": "acoes do banco do brasil (bbas3) tem avaliacao atraente apos 2021 forte, indica safra",
	"description": "de acordo com os analistas, papeis do banco do brasil (bbas3) tem recomendacao de compra, com preco-alvo a 50 por acao acoes do banco do brasil (bbas3) tem recomendacao de compra de r$ 50 por acao \ngetty images\na safra corretora aumentou o preco-alvo das acoe... ",
	"url": "https: //valor. Globo. Com/patrocinado/investe-safra/noticia/2022/03/07/acoes-do-banco-do-brasil-bbas3-tem-avaliacao-atraente-apos-2021-forte-indica-safra. Ghtml",
	"urltoimage": "https: //s2. Glbimg. Com/huvxx_6xu_dez9qexqgnuezuxxy=/1200x/smart/filters: cover(): strip_icc()/i. S3. Glbimg. Com/v1/auth_63b422c2caee4269b8b34177e8876b93/internal_photos/bs/2022/z/g/7q0uoeqzwa3d6p1jvphw/gettyimages-1203624927. Jpg",
	"publishedat": "2022-03-07t09:28:35z",
	"content": "acreditamos que o banco do brasil vem apresentando forte crescimento de credito (especialmente no credito rural e em alguns segmentos para pessoas fisicas) e boa qualidade da carteira e esta pronto p... [+4514 chars]"
	}
	]
	}`
