*********** сервис сокращения ссылок cutelinks https://github.com/CuteLinks/master запущен на cutelinks.ru ****************

C сохранением в postgres или redis.
Настройки находятся в config.json (строка подключения к бд, порт сервиса, тип хранилища).
Создан dockerfile, сервис сохранен в контейнере на dockerhub, можно скачать командой docker pull cutelinks/cutelinksgo

Пример для сокращения ссылки:
post cutelinks.ru:82
параметр формы url=fullurl.com
ответ - сокращенная ссылка dbLLQZoQPP или пустой ответ при повторяющемся url

Пример для получения исходной ссылки по сокращенной:
get cutelinks.ru:82/?url=dbLLQZoQPP
ответ - исходная ссылка fullurl.com или пустой ответ если ссылка не найдена

Имя домена в сокращенной ссылке не используется.
Будем считать что какой-то другой сервис будет подставлять например cutelinks.ru/dbLLQZoQPP и редиректить при необходимости на fullurl.com

Развёртывание сервиса:
Склонировать репо, распаковать бд, запустить docker-compose

git clone https://github.com/CuteLinks/master.git
cd master
tar -xzvf _data.tar.gz
docker-compose up --build app