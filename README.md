# Projeto Go + Elasticsearch + Kibana

Este projeto demonstra como integrar uma aplicação Go com Elasticsearch e Kibana usando Docker. O código foi refatorado para melhor legibilidade, separação de responsabilidades e configuração flexível via variável de ambiente.

## 🚀 Tecnologias Utilizadas
- **Go 1.21+**
- **Elasticsearch 8.10.2**
- **Kibana 8.10.2**
- **Docker & Docker Compose**

## 📦 Estrutura do Projeto
```
├── main.go               # Executável principal: conecta, testa e mostra exemplo de POI
├── docker-compose.yml    # Sobe Elasticsearch e Kibana
├── domain/
│   └── poi.go           # Estruturas POI e Location
├── elastic/
│   └── cliente.go       # Cliente Elasticsearch (configurável via env)
├── go.mod / go.sum       # Dependências Go
└── README.md             # Documentação
```

## ⚙️ Configuração Rápida

### 1. Clone o repositório
```bash
git clone https://github.com/SEU_USUARIO/elasticsearch-go-project.git
cd elasticsearch-go-project
```

### 2. Suba Elasticsearch e Kibana
```bash
docker-compose up -d
```

### 3. (Opcional) Configure o endereço do Elasticsearch
Por padrão, o endereço é `http://localhost:9200`. Para customizar, defina a variável de ambiente:
```bash
export ELASTICSEARCH_URL=http://seu-endereco:9200
```

### 4. Execute a aplicação Go
```bash
go run main.go
```

### 5. Acesse o Kibana
- [http://localhost:5601](http://localhost:5601)
- Dev Tools: [http://localhost:5601/app/dev_tools#/console](http://localhost:5601/app/dev_tools#/console)

## 📝 Exemplo de estrutura POI
```json
{
  "id": "1",
  "osm_id": "123",
  "type": "restaurante",
  "category_name": "food",
  "category_value": "italian",
  "name": "Pizzaria Central",
  "location": {
    "lat": -23.5505,
    "lon": -46.6333
  }
}
```

## 🔍 Exemplos de uso no Kibana Dev Tools

### Criar índice com mapping
```json
PUT /pois
{
  "mappings": {
    "properties": {
      "id": {"type": "keyword"},
      "osm_id": {"type": "keyword"},
      "type": {"type": "keyword"},
      "category_name": {"type": "keyword"},
      "category_value": {"type": "keyword"},
      "name": {"type": "text", "fields": {"keyword": {"type": "keyword"}}},
      "location": {"type": "geo_point"}
    }
  }
}
```

### Inserir documento
```json
POST /pois/_doc/1
{
  "id": "1",
  "osm_id": "123",
  "type": "restaurante",
  "category_name": "food",
  "category_value": "italian",
  "name": "Pizzaria Central",
  "location": {
    "lat": -23.5505,
    "lon": -46.6333
  }
}
```

### Buscar documentos
```json
GET /pois/_search
{
  "query": {
    "match": {
      "name": "Pizzaria Central"
    }
  }
}
```

## ✅ Funcionalidades
- Conexão e verificação de saúde do Elasticsearch
- Estruturas tipadas (POI, Location)
- Tratamento de erros robusto
- Configuração flexível via variável de ambiente
- Exemplo pronto para uso no Kibana
- Containerização completa

## 📚 Referências
- [Documentação oficial Elasticsearch Go](https://github.com/elastic/go-elasticsearch)
- [Documentação Kibana](https://www.elastic.co/guide/en/kibana/current/index.html)

---

> Código refatorado por boas práticas e pronto para evoluir!

{
  "mappings": {
    "properties": {
      "id": {"type": "keyword"},
      "osm_id": {"type": "keyword"},
      "type": {"type": "keyword"},
      "category_name": {"type": "keyword"},
      "category_value": {"type": "keyword"},
      "name": {
        "type": "text",
        "fields": {"keyword": {"type": "keyword"}}
      },
      "location": {"type": "geo_point"}
    }
  }
}
```

### Inserir documento:
```json
POST /pois/_doc/1
{
  "id": "1",
  "osm_id": "123",
  "type": "restaurante",
  "category_name": "food",
  "category_value": "italian",
  "name": "Pizzaria Central",
  "location": {
    "lat": -23.5505,
    "lon": -46.6333
  }
}
```

### Buscar documentos:
```json
GET /pois/_search
{
  "query": {
    "match": {
      "name": "Pizzaria Central"
    }
  }
}
```

## 🛠️ Funcionalidades

- ✅ Conexão com Elasticsearch
- ✅ Tratamento de erros (conexão, índice inexistente, documento não encontrado)
- ✅ Estruturas de dados tipadas (POI)
- ✅ Verificação de saúde do cluster
- ✅ Containerização com Docker

## 📦 Dependências

```bash
go get github.com/elastic/go-elasticsearch/v8
