https://doka.guide/recipes/github-new-pull-request/
https://youtu.be/zZBiln_2FhM?si=qUtp1THzBJ40Tb7u
https://youtu.be/VJm_AjiTEEc?si=jOKnMWAPVgyP89q6

Основные команды
**git init**
Позволяет проинициализировать репозиторий в текущей папке
**git status**
Показывает текущий статус
**git add**
Отслеживает изменения файлов git add index.html​ — добавляет index.html git add .​ — добавляет все файлы
**git commit**
Сохраняет изменения в коммит git commit -m 'commit message'​ — создает коммит с сообщением
**git branch**
Работа с ветками в репозитории git branch​ — показывает список веток git branch branch-name​ — создает новую ветку branch-name git branch -D branch-name​ — удаляет ветку branch-name
**git checkout**
Переключается на другую ветку git checkout branch-name​ — переключается на последний коммит в ветке branch-name
**git checkout -b branch-** name​ — создает и переключается на ветку branch-name
**git merge** Совмещает текущую ветку с выбранной git merge branch-name​ — совмещает текущую ветку с branch-name
**git config** Конфигурация и параметры git git config --global user.name​ — Показывает имя пользователя
**git config --global user.name 'new user'** ​ — Изменяет имя пользователя
**git config --global user.email** ​ — Показывает email пользователя
**git config --global user.email 'test@mail.ru**'​ — Изменяет email пользователя
**git push** Заливает текущие локальные коммиты в удаленный репозиторий
**git pull** Забирает изменения с удаленного репозитория в локальный
**git clone** Клонирует проект с удаленного репозитория