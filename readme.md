Telegram login bot
==================

Well, it's a telegram bot that notifies you if some login was made into a linux machine via ssh and it's made in Golang

How to use?
-----------

First of all you need to [create a new telegram bot](https://core.telegram.org/bots)
Once you have both a **token** and a **chat id** you should add the following line to `/etc/pam.d/sshd`

`session    optional     pam_exec.so /opt/loginbot -token=<token> -chatID=<chatID>`

Screenshot
----------
![alt text](screenshots/screenshot.png "A screenshot of the bot response in Telegram")