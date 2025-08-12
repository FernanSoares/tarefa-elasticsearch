package elastic

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewElasticClient cria um cliente Elasticsearch, lendo o endereço do servidor da variável de ambiente ELASTICSEARCH_URL se disponível.
func NewElasticClient(user, pass string) (*elasticsearch.Client, error) {
	addr := os.Getenv("ELASTICSEARCH_URL")
	if addr == "" {
		addr = "http://localhost:9200"
	}
	cfg := elasticsearch.Config{
		Addresses: []string{addr},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cliente Elasticsearch: %w", err)
	}

	res, err := client.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro de conexão com Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch não está respondendo (Status: %s)", res.Status())
	}

	return client, nil
}

// Função adicional para verificar saúde do cluster
func CheckClusterHealth(client *elasticsearch.Client) error {
	res, err := client.Cluster.Health()
	if err != nil {
		return fmt.Errorf("erro ao verificar saúde do cluster: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("cluster Elasticsearch com problemas (Status: %s)", res.Status())
	}

	return nil
}
