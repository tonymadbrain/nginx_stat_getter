# Nginx stat getter for Zabbix monitoring

# ENG

Script for getting statistics from [Nginx](https://nginx.org/) Web server for Zabbix external check. Implementation of this script - https://github.com/vicendominguez/nginx-zabbix-template on Golang.

## Preparation
### Nginx

To enable statistics in Nginx, create `virtual host`, like so

```Bash
server {
  listen 4040; # unusual port for monitoring
  server_name _; # doesn't care about server name
  keepalive_timeout 0;
  allow 192.168.0.40; # allow requests only from our monitoring server
  deny all; # deny all others
  location =/nginx_status/ {
    stub_status on; # properly enabling of statistics
  }
  access_log off; # don't write logs
}
```

And reload/restart the nginx service.

### Zabbix

We need to compile binary for that platform where zabbix is running. Use this command:

````Bash
env GOOS={OS} GOARCH={ARCH} go build -v github.com/tonymadbrain/nginx_stat_getter
````

Where:

{OS} - os type:

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

{ARCH} - arhitecture:

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

Then, copy the binary to Zabbix server into `/usr/lib/zabbix/externalscripts` folder, make him executable with `chmod +x nginx_stat_getter`, set zabbix owner with  `chown zabbix:zabbix nginx_stat_getter`.
Next, import template `zbx_nginx_template.xml` in Zabbix frontend and attach him to server(s).

Done! Wait for data.

# RUS

Скрипт для получения статистики из Nginx, который можно использовать в Zabbix. Реализация на Go вот этого скрипта - https://github.com/vicendominguez/nginx-zabbix-template.

## Настройка
### Nginx

Чтобы включить статистику в Nginx, нужно создать vhost, следующего содержания:

```Bash
server {
  listen 4040; # нестандартный порт для мониторинга статистики
  server_name _; # нам пофиг на server name
  keepalive_timeout 0;
  allow 192.168.0.40; # разрешаем запросы только для ip адреса нашего сервера мониторинга
  deny all; # все остальные идут лесом
  location =/nginx_status/ {
    stub_status on; # собственно включение статистики
  }
  access_log off; # не пишем логи
}
```

И сделать reload/restart сервера.

### Zabbix

Нужно скомпилировать бинарник под ту платформу, на которой запущен Zabbix сервер, для этого нужно использовать команду:

````Bash
env GOOS={OS} GOARCH={ARCH} go build -v github.com/tonymadbrain/nginx_stat_getter
````

{OS} - тип операционной системы, может быть:

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

{ARCH} - архитектура, может быть:

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

Закинуть бинарник на сервер Zabbix в каталог `/usr/lib/zabbix/externalscripts`, сделать его исполняемым - `chmod +x nginx_stat_getter`, сделать владельцем файла Zabbix - `chown zabbix:zabbix nginx_stat_getter`. Затем нужно импортировать шаблон `zbx_nginx_template.xml` в Zabbix фронтенде и прикрепить его к нужному серверу.

Ждать данных.
