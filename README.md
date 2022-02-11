# GO CLI- client and aiohttp server 

Сервер: python aiohttp, 5 API.

    - GET, получить список файлов и их hash по содержимому на сервере в папке /tmp

    - GET, получить по имени файла из папки /tmp файл, если файла нет - возвращать ошибку

    - PUT, положить файл в папку /tmp, если уже файл есть - возвращать ошибку

    - POST, обновить файл в папке /tmp новым файлом, если файла нет - возвращать ошибку, если файл есть 
        и по hash совпадает -возвращать что не требуется обновление

    - DELETE, удалить по имени файла из папки /tmp файл, если файла нет - возвращать ошибку

Собрать в Docker. Командой через docker-compose должен подниматься контейнер, 
к которому можно обращаться по API.

Клиент: Go CLI-приложение. Скомпилировать в исполняемое приложение, которое работает 
по параметрам. Рядом с приложением - conf файл, в котором указаны адрес и порт сервера. 
Предусмотреть вариант передачи адреса и порта через командную строку.

Запуск приложения:

    - Получить список файлов. Обращается к серверу, получает список файлов, выводит на экран список

    - Прочитать файл с сервера и записать на клиента. Обращается к серверу, получает 
        файл по имени, записывает по указанному в параметрах пути, выводит результат: Ок/НеОк

    - Записать на сервер файл. Обращается к серверу, отдает файл на запись, 
        выводит результат Ок/НеОк

    - Обновить на сервере файл. Обращается к серверу, отдает файл на перезапись, 
        выводит результат Ок/НеОк

    - Удалить на сервере файл. Обращается к серверу, выводит результат Ок/НеОк


Инструкция по запуску приложения:
1. Склонируйте репозиторий
2. В папке aiohttp_files_hash/ выполнить комманды: 
   1. docker-compose build
   2. docker-compose up
    запуститься сервер по адресу: http://0.0.0.0:8080 
3. В папке go_client выполнить комманды:
   1. go build
   2. ./go_client
4. следовать запросам появляющимся на экране.

- в папке go_client находися файл "config.cfg" в котором находятся настройки для работы 
  приложения:
  - первая строка адресс сервера в виде: http://0.0.0.0
  - второя строка порт сервера: 8080
  - третья строка путь для сохранения файла: /folder1/folder2/folder3/folder4/
- при выборе опции ввода конфигурации с клавиатуры формат вводимых данных должен
  соответсвовать формату данных проведенных выше.