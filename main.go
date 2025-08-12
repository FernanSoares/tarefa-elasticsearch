package main

import (
	"encoding/json"
	"fmt"
	"log"

	"elastic-project/domain"
	"elastic-project/elastic"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	client := connectElasticOrExit()
	defer clientCloseInfo(client)
	printStatusMessages()
	showPOIExample()
}

func connectElasticOrExit() *elasticsearch.Client {
	es, err := elastic.NewElasticClient("", "")
	if err != nil {
		log.Fatalf("Erro ao criar cliente: %v", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Erro de conexão: %v", err)
	}
	if res.IsError() {
		log.Fatalf("Erro na resposta: %s", res.String())
	}
	return es
}

func clientCloseInfo(es *elasticsearch.Client) {
	// Não é necessário fechar o cliente, mas pode-se adicionar logs ou métricas aqui futuramente
}

func printStatusMessages() {
	fmt.Println("✅ Conexão com Elasticsearch estabelecida!")
	fmt.Println("🔗 Elasticsearch está rodando em: http://localhost:9200")
	fmt.Println("🔗 Kibana está disponível em: http://localhost:5601")
}

func showPOIExample() {
	poi := domain.POI{
		ID:            "1",
		OSMID:         "123",
		Type:          "restaurante",
		CategoryName:  "food",
		CategoryValue: "italian",
		Name:          "Pizzaria Central",
		Location:      domain.Location{Lat: -23.5505, Lon: -46.6333},
	}

	data, _ := json.MarshalIndent(poi, "", "  ")

	fmt.Println("\n📋 Exemplo de estrutura POI em JSON:")
	fmt.Println(string(data))
	fmt.Println("\n📝 Use esse formato no Kibana Dev Tools para inserir dados!")
}
