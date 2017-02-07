# Nginx stat getter for Zabbix monitoring

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
env GOOS={OS} GOARCH={ARCH} go build -v github.com/username/nginx_stat_getter
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
