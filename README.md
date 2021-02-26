# Slashy
A test of Discord's new slash commands and interaction webhooks, built in Golang.

You can try Slashy by [adding the slash commands to a server](https://go.maia.codes/slashy-invite), or by [joining the Slashy Discord server](https://discord.gg/YeZWFWkxvT).

<a href="https://chat.assistcord.com/#slashy" target="_blank"><img height="40px" src="https://cdn.assistcord.com/button/chat_with_us.svg"></img></a>

## How it works
Slashy uses Discord's new Bot Interaction Endpoint for slash commands, meaning Discord sends a web request to a web server within Slashy for every interaction event.

## Why?
Discord recently introduced the slash commands API, but the only examples I could find were with JS and Python, So I decided that I'd create my own slash commands bot from scratch in Golang, as it's my favourite language.

## Running Slashy yourself
To use Slashy you'll need to use Go to build into a binary (I'm just going to expect you know how to do this already but join the server if you need help), Then you'll need pass environment variables to Slashy. Currently it doesn't support use of a `.env` file so you'll have to do this in bash. I recommend making an `.sh` file and `export`ing the variables you need before running the binary. Here are the environment variables you'll need:
- `client_id` - your Discord app's Client ID.
- `development` - `true`/`false` whether the bot should be run in development mode.
- `public_key` - The public key found in your Discord developer dashboard.
- `bot_token` - Currently Discord requires a bot token to register commands, this will be changed in the future.
- `PORT` - Which port to run the web server on.

## How Slashy runs in Production
Slashy currently runs on Google Cloud Run, this is useful because it allows Slashy to instantly scale to demand without any infrastructure management, as Cloud Run is serverless. This repo has Google Cloud Build setup to automatically deploy updates to production.

## Anonymous Statistics
Slashy now collects statistics which includes:
- The Guild ID in which commands are run
- Which commands are being run
- When commands are being run
