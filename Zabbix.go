package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Estrutura para o JSON de resposta
type Host struct {
	ID   string `json:"hostid"`
	Nome string `json:"host"`
}

type Resposta struct {
	Resultados []Host `json:"result"`
}

func main() {
	// URL da API Zabbix
	url := "http://-----//zabbix/api_jsonrpc.php"

	// Token de autenticação
	tokenDeAutenticacao := "---"

	// Cabeçalho da requisição
	cabecalhos := map[string]string{
		"Content-Type": "application/json-rpc",
	}

	pedido := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "host.get",
		"params": map[string]interface{}{
			"output": []string{"hostid", "host"},
		},
		"auth": tokenDeAutenticacao,
		"id":   1,
	}

	pedidoBytes, erro := json.Marshal(pedido)
	if erro != nil {
		log.Fatalf("Erro ao criar o pedido JSON: %v", erro)
	}

	requisicao, erro := http.NewRequest("POST", url, bytes.NewBuffer(pedidoBytes))
	if erro != nil {
		log.Fatalf("Erro ao criar a requisição: %v", erro)
	}

	for chave, valor := range cabecalhos {
		requisicao.Header.Set(chave, valor)
	}

	client := &http.Client{}
	resposta, erro := client.Do(requisicao)
	if erro != nil {
		log.Fatalf("Erro na requisição: %v", erro)
	}
	defer resposta.Body.Close()

	if resposta.StatusCode != 200 {
		log.Fatalf("Erro na requisição, código de status: %d", resposta.StatusCode)
	}

	var respostaFinal Resposta
	erro = json.NewDecoder(resposta.Body).Decode(&respostaFinal)
	if erro != nil {
		log.Fatalf("Erro ao decodificar a resposta JSON: %v", erro)
	}

	if len(respostaFinal.Resultados) == 0 {
		fmt.Println("Nenhum host encontrado.")
		return
	}

	// Criação do arquivo CSV
	arquivo, erro := os.Create("Relatorio_Hosts_.csv")
	if erro != nil {
		log.Fatalf("Erro ao criar o arquivo CSV: %v", erro)
	}
	defer arquivo.Close()

	// Configurando o writer CSV para usar ponto e vírgula como delimitador
	writer := csv.NewWriter(arquivo)
	writer.Comma = ';' // Define o delimitador como ponto e vírgula
	defer writer.Flush()

	// Adicionando cabeçalhos ao CSV
	cabecalhosCSV := []string{"ID do Host", "Nome do Host"}
	erro = writer.Write(cabecalhosCSV)
	if erro != nil {
		log.Fatalf("Erro ao escrever no arquivo CSV: %v", erro)
	}

	// Escrevendo os dados dos hosts no CSV
	for _, host := range respostaFinal.Resultados {
		erro := writer.Write([]string{host.ID, host.Nome})
		if erro != nil {
			log.Fatalf("Erro ao escrever no arquivo CSV: %v", erro)
		}
	}

	// Mensagem indicando sucesso
	fmt.Println("Relatório exportado com sucesso para 'Relatorio_Hosts_.csv'.")
}
