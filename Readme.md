# VK User Info

This is a small pet project written in Go programming language. It loads basic information of user by ID and calculates its age.

## Compilation

```bash
go build -o vk_user_info main.go vk.go
```

## Usage

```bash
./vk_user_info -t *vk_auth_token* *vk_user_id*
```

For example:
```
./vk_user_info -t *vk_auth_token* durov
Getting general available information about "durov"...
Calculating user's age...

==========
USER DATA:
==========
ID: 1
Name: Павел Дуров
Nickname:
Sex: Man
Bdate: 10.10.1984 (34 years)
City: Санкт-Петербург
Country: Россия
Relation: 0
==========

```

## How to get VK Auth Token

To get VK authentication token proceed to [VK authentication link](https://oauth.vk.com/authorize?client_id=6826999&scope=&redirect_uri=https://oauth.vk.com/blank.html&display=wap&v=5.63&response_type=token), pass the login procedure and copy `access_token` parameter from URL.
