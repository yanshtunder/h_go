## ДЗ 1
### Задание
Реализовать утилиту **sort**.

Использовать пакет `flag`, пакет `slices`, пакет `strings`, пакет `unicode/utf8`

Отсортировать строки в файле по аналогии с консольной утилитой sort : на входе
подается файл из несортированными строками, на выходе - файл с отсортированными.

**Реализовать поддержку утилитой следующих ключей:**

1. [x] `-f` - путь к файлу, который необходимо отсортировать
2. [x] `-k` - указание колонки для сортировки (слова в строке могут выступать в 
качестве колонок, по умолчанию разделитель - пробел)
3. [ ] `-n` - сортировать по числовому значению
4. [x] `-r` - сортировать в обратном порядке
5. [x] `-u` - не выводить повторяющиеся строки

**Дополнительно 1**

1. [ ] `-M` - сортировать по названию месяца
2. [ ] `-b` - игнорировать хвостовые пробелы
3. [ ] `-c` — проверять отсортированы ли данные
4. [ ] `-h` — сортировать по числовому значению с учетом суффиксов

**Долнительно 2** 

1. [ ] Сделать unit-tests на весь функционал

<hr>

### Запуск

``` bash
cd bin
./sort -f=../test.txt -r -u
./sort -f=../input.txt -k=1 -r
cd ../
```

