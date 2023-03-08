# gpt-cli

gpt-cli is a command-line interface tool for OpenAI's GPT-3 language model using the GPT3 API. It can generate text based on the input prompt provided by the user.

## Installation

To install GPT-CLI, the user must have python and pip already installed. Run the following command in the terminal:

```sh
go build -o gpt-cli
```

## Usage

After installing GPT-CLI, the user can access it through the command-line interface. To generate text, the user must provide their GPT-3 API key, the prompt, and the length of the output text.
If you are using it for the first time, you will be asked to enter your OpenAI API key. The key will be stored on your drive and will not be necessary for further uses.

```sh
./gpt-cli 'write me quote of the day'
```

## Limitations

GPT-CLI is currently limited to generating only 2048 tokens per request due to the restriction of the GPT-3 API. In addition, the tool is still in development and may contain bugs.

## License

GPT-CLI is licensed under the MIT license.

## Contributions

Contributions to GPT-CLI are welcome. Please fork the repository and submit a pull request.
