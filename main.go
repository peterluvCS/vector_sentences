package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/tmc/langchaingo"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	connStr := "postgres://laotang:tanghan123456@localhost:5432/vector_tb"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Query the data from PgVector table
	rows, err := conn.Query(context.Background(), "SELECT embedding FROM vector_tb")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	var vectors []string
	for rows.Next() {
		var vector string
		if err := rows.Scan(&vector); err != nil {
			log.Fatalf("Row scan failed: %v\n", err)
		}
		vectors = append(vectors, vector)
	}

	if rows.Err() != nil {
		log.Fatalf("Row iteration failed: %v\n", rows.Err())
	}

	// Initialize Ollama client
	ollamaAPIKey := os.Getenv("8186890a1eacea4d9c46c6f8468325c5a516ca5b896abd4b44c1c00aa34bcff0")
	client, err := ollama.NewClient(ollamaAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Ollama client: %v\n", err)
	}

	// Convert vectors to meaningful sentences using LangChain and Ollama
	for _, vector := range vectors {
		prompt := fmt.Sprintf("Convert the following vector to a meaningful sentence: %s", vector)

		response, err := client.Complete(context.Background(), langchain.CompletionRequest{
			Prompt:    prompt,
			MaxTokens: 100,
		})
		if err != nil {
			log.Fatalf("Failed to get completion: %v\n", err)
		}

		if len(response.Choices) > 0 {
			fmt.Printf("Vector: %s\nSentence: %s\n", vector, response.Choices[0].Text)
		} else {
			fmt.Printf("Vector: %s\nSentence: No response\n", vector)
		}
	}
}
