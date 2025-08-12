# 🔍 Elasticsearch POI Search Engine

Uma aplicação moderna em Go para indexação e busca de Pontos de Interesse (POI) utilizando Elasticsearch como motor de busca e Kibana para visualização de dados.

---

## 📋 Sobre o Projeto

Esta aplicação foi desenvolvida para demonstrar a integração eficiente entre Go e Elasticsearch, focando em operações de busca geoespacial e indexação de dados estruturados. O sistema permite armazenar e consultar informações sobre locais de interesse com coordenadas geográficas.

## 🛠️ Stack Tecnológica

| Tecnologia | Versão | Propósito |
|------------|--------|-----------|
| Go | 1.21+ | Backend e lógica de negócio |
| Elasticsearch | 8.10.2 | Motor de busca e indexação |
| Kibana | 8.10.2 | Interface de visualização |
| Docker Compose | - | Orquestração de containers |

## 🏗️ Arquitetura

```
tarefa-elasticsearch/
│
├── 🎯 main.go                    → Ponto de entrada da aplicação
├── 🐳 docker-compose.yml         → Definição dos serviços
├── 📁 domain/
│   └── poi.go                    → Modelos de dados (POI/Location)
├── 📁 elastic/
│   └── cliente.go                → Abstração do cliente Elasticsearch
├── 📄 go.mod                     → Gerenciamento de dependências
└── 📄 go.sum                     → Checksums das dependências
```

## 🚀 Guia de Execução

**Pré-requisitos:**
- Docker e Docker Compose instalados
- Go 1.21 ou superior

**Passos para executar:**

1. **Inicializar os serviços:**
   ```bash
   docker-compose up -d
   ```

2. **Aguardar inicialização completa:**
   ```bash
   # Verificar se Elasticsearch está respondendo
   curl -X GET "localhost:9200/_cluster/health?pretty"
   ```

3. **Executar a aplicação Go:**
   ```bash
   go run main.go
   ```

## 🌐 Configuração de Ambiente

A aplicação suporta configuração via variáveis de ambiente:

```bash
# URL do cluster Elasticsearch (padrão: http://localhost:9200)
export ELASTICSEARCH_URL=http://localhost:9200
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
