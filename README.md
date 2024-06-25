# vector_sentences

## Prerequisites

- Go 1.20 or later
- PostgreSQL database with PgVector extension
- Ollama API key

## Flow of the Application

1. **Database Connection**:
    - Connects to the PostgreSQL database using the provided connection string.
    - Queries the vectors from the specified table.

   **Sample Output**:
    ```
    Connected to the database successfully.
    Fetched vectors from the database.
    ```

2. **Initializing Ollama Client**:
    - Creates an Ollama client using the API key.

   **Sample Output**:
    ```
    Ollama client initialized successfully.
    ```

3. **Converting Vectors to Sentences**:
    - Iterates through each vector.
    - Sends a prompt to the Ollama API to convert the vector into a meaningful sentence.
    - Prints the converted sentence.

   **Sample Output**:
    ```
    Vector: [0.1, 0.2, 0.3, ...]
    Sentence: This is a meaningful sentence generated from the vector.
    ```

