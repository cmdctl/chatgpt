# Golang OpenAI Chatbot

This project is a simple Golang chatbot that uses OpenAI's API to generate responses. It reads input from stdin and returns a response based on the provided messages. 

## Prerequisites

- [Go](https://golang.org/doc/install)

## Installation

1. Clone the repository:

```
git clone https://github.com/username/go-openai-chatbot.git
```

2. Install dependencies:

```
go mod download
```

3. Set up an OpenAI API key with the `openai` package installed. You can sign up for a free API key [here](https://beta.openai.com/signup/).

4. Set your OpenAI API key as an environment variable:

```
export OPENAI_KEY=your-api-key
```

## Usage

The chatbot reads input from stdin and uses the provided delimiter to split the inputted messages into chunks. It then sends these chunks as messages to the OpenAI API to generate a response.

```
go run main.go -d "--------"
```

For example:

```
echo "Hello, how are you today?--------I'm doing well. How about you?" | go run main.go -d "--------"
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.


