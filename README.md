<h1 align="center">Gobot : Messenger ChatBot Golang</h1>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/hbollon/gobot" target="_blank">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/hbollon/gobot" />
  </a>
  <a href="https://github.com/hbollon/gobot/blob/master/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> Facebook Messenger chatbot using the Levenshtein distance algorithm for pattern matching. Use mux and yaml.v2.

---

## Table of Contents

- [Requirements](#requirements)
- [Presentation](#presentation)
- [Features](#features)
- [Installation](#installation)
- [Customize chatbot responses](#customize-chatbot-responses)
- [Author](#author)
- [Contributing](#-contributing)
- [License](#-license)


---

## Requirements
- [Go](https://golang.org/doc/install)
- Facebook account

## Presentation
This chatbot was developed in order to manage efficiently any facebook page. It use Facebook API.
Gobot use Levenshtein distance algorithm for pattern matching to calculate a match percentage between the received message and templates defined in content.yml in order to be more flexible for the interlocutor instead of only responding to messages that match exactly.
For example, if you have set a minimal matching percentage of 40 % and defined patterns like this in content.yml :

```yaml
templates: 
- messages: ["Hello", "Hi"]
response: "Hi !\nI'm GoBot chat bot ! :)"
- messages: ["Who's your creator ?"]
response: "I was build by @hbollon from Bits Please Inc. :)\nMy source code is available on his Github : github.com/hbollon !"

default_response: "Sorry, I don't understand your message... :/\nPlease try again with different sentence or using more words."
```

And someone sends you "Who's build you ?"
It will match with "Who's your creator ?" and send back corresponding response !
Same with "Creator" :)

However, if it doesn't match with any template, it will send the **default_response**.
You can edit the minimal matching percentage in the config.yml file.

## Features
- Designed for Facebook Messenger ✨
- Easily configurable through yaml files ✨
- Efficient pattern matching with Levenshtein distance ✨
- Customizable minimal matching percentage ✨

## Installation
### Clone
- Firstly, clone this repo wherever you want using :
```git 
git clone https://github.com/hbollon/gobot
```

### Setup
- Create an app on Facebook developer : https://developers.facebook.com/apps/

- Add it Messenger product and configure your webhook by linking a Facebook page :
<img align="center" src="https://i.ibb.co/VWzZNZh/page.png" alt="page" border="0" />

- Configure your webhook Callback URL (You can use [Ngrok](https://ngrok.com/) to make public URL to your localhost for testing purpose), the Verify Token and subscribe to messages field:
<img align="center" src="https://i.ibb.co/NLHNFTB/webhook.png" alt="webhook" border="0" />

- Edit configs/config.yml file with your Messenger App credentials : 
```yaml
# Facebook API config
app_secret: <app_secret> # Secret key which can be found in the app settings on Facebook Developer dashboard
access_token: <facebook_app_token> # Generated token on app Messenger product
verify_token: <webhook_verify_key> # Verify token of your webhook URL 
```
You can copy config.yml.exemple to config.yml and replace corresponding values.

- Open bash in root project directory and run :
```bash
go run cmd/gobot/*.go # Run program without build it
# or
go build cmd/gobot/*.go && ./gobot # Build and run

```

## Customize chatbot responses
In order to modify the messages that your chatbot will recognize as well as the responses that it will send back, you just need to modify the content.yml file in the config folder.

Exemple:
```yaml
templates: 
    - messages: ["Hello", "Hi"]
      response: "Hi !\nI'm GoBot chat bot ! :)"
    - messages: ["Who's your creator ?"]
      response: "I was build by @hbollon from Bits Please Inc. :)\nMy source code is available on his Github : github.com/hbollon !"

default_response: "Sorry, I don't understand your message... :/\nPlease try again with different sentence or using more words."
```

A template is composed of **messages** which are the sentences on which the bot is based to try to understand the interlocutor via the Levenshtein algorithm. If the received message matches one of these messages (depending on the match threshold defined in config.yml) then the response contained in the **response** attribute will be sent to the user.
If received message don't match with any template, the bot will send back the **default_response**.

## Author

👤 **Hugo Bollon**

* Github: [@hbollon](https://github.com/hbollon)
* LinkedIn: [@Hugo Bollon](https://www.linkedin.com/in/hugo-bollon-68a2381a4/)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/hbollon/gobot/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Hugo Bollon](https://github.com/hbollon).<br />
This project is [MIT License](https://github.com/hbollon/gobot/blob/master/LICENSE.md) licensed.
