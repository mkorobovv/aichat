# AI Chat

A small demo project showcasing how to integrate with LLM (Large Language Model) APIs.
The goal is to demonstrate a simple chat interface and project structure for working with language models.

## Features
- Send messages to an LLM via API
- Receive chat-style responses
- Keep in memory chat history

## Installation & Setup

1. Clone the repository:

```shell
git clone https://github.com/mkorobovv/aichat.git
cd aichat
```

2. Setup environment variables.

3. Run with Docker Compose.

```shell
docker-compose up --build
```

## API documentation

&emsp;`http://localhost:8080/api/v1/send` 

**METHOD:** POST

**Request body**

```json5
{
    "user_id": 12,
    "chat_id": 8,
    "content": "What is GPT?"
}
```

**Response body**

```json5

```