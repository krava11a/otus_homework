Пометки из лекции и ДЗ.

docker run  -t -i tarantool/tarantool -i

lua & sql

box.schema.space.create('dialogs') - создание space
далее индекс
box.space.dialogs:create_index('primary', {type="TREE", unique=true,parts={1,'unsigned'}})
insert:
box.space.dialogs:insert({1,'werty',6})
box.space.dialogs:insert({1,23,'werty',6})
select (боль и страдания):
box.space.dialogs:select(2) - select по ключу работает
box.space.dialogs:select({1},{iterator='GT'}) - select с условием больше
update:
замена второго значения в тапле [2, 23, 'werty', 6] на 45 будеи выглядеть так:
box.space.dialogs:update({2},{{'=',2,45}}),
плюс в том что можно использовать разные операторы +,- и другие.

Хранимые процедуры.

Можно доставить модули. Модули ставяться через rocks.ctl ?
Ставила через require в lua.

Отдельный гемор это декодировать полученные значения из тарантул и перекладывать их в требуемые по интерфейсу.