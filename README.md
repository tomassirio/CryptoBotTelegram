[comment]: <> (<p align="center"><img src="https://i.imgur.com/a1H1sQa.png"/></p>)

[comment]: <> (<p align="center">)

[comment]: <> (  <h1 align="center">BITCOIN-TELEGRAM-BOT</h1>)

[comment]: <> (  <p align="center">)

[comment]: <> (    ·)

[comment]: <> (    <a href="https://github.com/tomassirio/bitcoinTelegramBot/issues">Report Bug</a>)

[comment]: <> (    ·)

[comment]: <> (    <a href="https://github.com/tomassirio/bitcoinTelegramBot/issues">Request Feature</a>)

[comment]: <> (    ·)

[comment]: <> (  </p>)

[comment]: <> (</p>)

[comment]: <> (<p align="center">)

[comment]: <> (  <a href="https://github.com/tomassirio/bitcoinTelegramBot/graphs/contributors"><img src="	https://img.shields.io/github/issues/tomassirio/BitcoinTelegramBot"></a>)

[comment]: <> (  <a href="https://github.com/tomassirio/BitcoinTelegramBot/blob/master/LICENSE"><img src="https://img.shields.io/github/license/tomassirio/BitcoinTelegramBot"></a>)

[comment]: <> (  <a href="https://github.com/tomassirio/bitcoinTelegramBot/network/members"><img src="https://img.shields.io/github/forks/tomassirio/BitcoinTelegramBot"></a>)

[comment]: <> (  <a href="https://img.shields.io/github/stars/tomassirio/bitcoinTelegramBot"><img src="https://img.shields.io/github/stars/tomassirio/BitcoinTelegramBot"></a>)

[comment]: <> (</p>)

[comment]: <> (BitcoinTelegram is a simple Telegram Bot to consult BTC's price over Telegram. It's written in Golang with Telebot's Framework)

[comment]: <> (<!-- TABLE OF CONTENTS -->)

[comment]: <> (<details open="open">)

[comment]: <> (  <summary>Table of Contents</summary>)

[comment]: <> (  <ol>)

[comment]: <> (    <li>)

[comment]: <> (      <a href="#tada-inspiration">Inspiration</a>)

[comment]: <> (    </li>)

[comment]: <> (    <li>)

[comment]: <> (      <a href="#star-getting-started">Getting Started</a>)

[comment]: <> (      <ul>)

[comment]: <> (        <li><a href="#what-you-will-need">What you will need</a></li>)

[comment]: <> (        <li><a href="#computer-installation">Installation</a></li>)

[comment]: <> (        <li><a href="#white_check_mark-add-btb-to-your-telegram-channel">Add BTB to your Telegram Channel</a></li>)

[comment]: <> (      </ul>)

[comment]: <> (    </li>)

[comment]: <> (    <li><a href="#battery-usage">Usage</a></li>)

[comment]: <> (    <li><a href="#building_construction-contribution-guidelines">Contribution Guidelines</a></li>)

[comment]: <> (    <li><a href="bust_in_silhouette-who-am-i">Who Am I?</a></li>)

[comment]: <> (  </ol>)

[comment]: <> (</details>)

[comment]: <> (## :tada: Inspiration)

[comment]: <> (Over the last couple of months I grew fond of Golang, Telegram and Bitcoin. So I decided to combine those three elements to create something that could get a little bit of each world. Hence, Bitcoin-Telegram-Bot &#40;in Golang&#41; was created)

[comment]: <> (## :star: Getting started)

[comment]: <> (### What you will need:)

[comment]: <> (- You are going to need a computer or server where to host the bot.)

[comment]: <> (- Git)

[comment]: <> (- Golang v1.13)

[comment]: <> (- A device with Telegram)

[comment]: <> (### :computer: Installation)

[comment]: <> (Open a Terminal and copy these commands &#40;Linux & Mac devices&#41;:)

[comment]: <> (```bash)

[comment]: <> (cd ~)

[comment]: <> (git clone https://github.com/tomassirio/BitcoinTelegramBot.git)

[comment]: <> (cd ./BitcoinTelegramBot)

[comment]: <> (mv .env.example .env)

[comment]: <> (go get github.com/tomassirio/bitcoinTelegram)

[comment]: <> (go run main.go)

[comment]: <> (```)

[comment]: <> (##### Warning: )

[comment]: <> (This won't work unless you replace the **REPLACE_WITH_TOKEN** on the .env file with the Token granted by @BotFather)

[comment]: <> (### :white_check_mark: Add BTB to your Telegram Channel)

[comment]: <> (Open [@BotFather]&#40;https://telegram.me/botfather&#41; on telegram and create a new bot with it's __/newbot__ command.)

[comment]: <> (Assign it a name. This name won't be the one that is shown on each message, so you can name it whatever you want.)

[comment]: <> (@BotFather will grant you a Token. This token is the one that will replace the **REPLACE_WITH_TOKEN** on the .env.example file on this repository. &#40;Don't forget to rename that file to .env&#41;)

[comment]: <> (![image]&#40;https://i.imgur.com/RC8anHA.png&#41;)

[comment]: <> (You can also play a little bit more with @BotFather. For example you can use the __/setcommands__ to define the uses your bot has on the '/' icon:)

[comment]: <> (```)

[comment]: <> (price - Gets BTC actual price)

[comment]: <> (historic - Gets a percentage between Today's and Yesterday's price)

[comment]: <> (summary - Gets both the price and historic values)

[comment]: <> (```)

[comment]: <> (![image]&#40;https://i.imgur.com/ACmSAF1.png&#41;)

[comment]: <> (## :battery: Usage)

[comment]: <> (Once the bot is running and added to your Telegram Group, you can use any of the following commands:)

[comment]: <> (```sh)

[comment]: <> (    * /price : Get's bitcoin's Last price)

[comment]: <> (    * /historic : Gets a percentage between Today's and Yesterday's price)

[comment]: <> (    * /summary : Gets both the price and historic values)

[comment]: <> (```)

[comment]: <> (## :building_construction: Contribution Guidelines:)

[comment]: <> (-   **_fork_** and **_clone_** this repository)

[comment]: <> (-   Make a new branch using `git checkout -b change/username`)

[comment]: <> (-   Commit the desired changes to that branch)

[comment]: <> (-   Sign off your commits using `git commit -s -m w/signoff`)

[comment]: <> (-   Push your changes to the branch and open a pull request)

[comment]: <> ( -->)

[comment]: <> (## :bust_in_silhouette: Who Am I?)

[comment]: <> (<img src="https://media.discordapp.net/attachments/763140054825697301/763681938652528690/logo-design-branding-logo-tool-open-electronic-1-5f7ed02bc8247.png?width=468&height=468" width="410" height="410" /></p>)

[comment]: <> (  <a href="mailto:tomassirio@gmail.com?Subject=Tomas%20You%20Are%20Amazing!">)

[comment]: <> (      <img src="https://cdn2.downdetector.com/static/uploads/logo/image21.png" width="100"; height="100"/>)

[comment]: <> (  </a>)

[comment]: <> (  <a href="https://www.linkedin.com/in/tomassirio/">)

[comment]: <> (      <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimage.flaticon.com%2Ficons%2Fpng%2F512%2F174%2F174857.png&f=1&nofb=1" width="100"; height="100"/>)

[comment]: <> (  </a>)

[comment]: <> (  <a href="https://dev.to/tomassirio">)

[comment]: <> (      <img src="https://avatars3.githubusercontent.com/u/13521919?s=280&v=4" width="100"; height="100"/>)

[comment]: <> (  </a>)

[comment]: <> (  <a href="https://www.buymeacoffee.com/tomassirio1">)

[comment]: <> (      <img src="https://i.pinimg.com/originals/60/fd/e8/60fde811b6be57094e0abc69d9c2622a.jpg" width="100"; height="100"/>)

[comment]: <> (  </a>)