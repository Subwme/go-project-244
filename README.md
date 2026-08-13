### Hexlet tests and linter status:
[![Actions Status](https://github.com/Subwme/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Subwme/go-project-244/actions)
[![CI](https://github.com/Subwme/go-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/Subwme/go-project-244/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Subwme_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Subwme_go-project-244)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Subwme_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Subwme_go-project-244)

## gendiff

Утилита командной строки для сравнения двух конфигурационных файлов (JSON или YAML) и отображения их различий.

## Установка

```bash
git clone https://github.com/Subwme/go-project-244.git
cd go-project-244
make build
```

## Использование

```
./bin/gendiff [global options] <filepath1> <filepath2>

GLOBAL OPTIONS:
   --format format, -f format  output format [stylish|plain|json] (default: "stylish")
   --help, -h                  show help
```

## Форматы вывода

### stylish (по умолчанию)

Отображает дерево различий с отступами и знаками `+`/`-`:

```bash
./bin/gendiff testdata/fixtures/file1.json testdata/fixtures/file2.json
```

```
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

### plain

Описывает изменения в виде читаемых предложений:

```bash
./bin/gendiff --format plain testdata/fixtures/file1.json testdata/fixtures/file2.json
```

```
Property 'follow' was removed
Property 'proxy' was removed
Property 'timeout' was updated. From 50 to 20
Property 'verbose' was added with value: true
```

### json

Выдаёт структурированный JSON-объект с типом каждого изменения:

```bash
./bin/gendiff --format json testdata/fixtures/file1.json testdata/fixtures/file2.json
```

```json
{
    "follow": {
        "type": "removed",
        "value": false
    },
    "host": {
        "type": "unchanged",
        "value": "hexlet.io"
    },
    "proxy": {
        "type": "removed",
        "value": "123.234.53.22"
    },
    "timeout": {
        "new": 20,
        "old": 50,
        "type": "changed"
    },
    "verbose": {
        "type": "added",
        "value": true
    }
}
```

## Demo

[![asciicast](https://asciinema.org/a/r912nFRzg0gzmW0J.svg)](https://asciinema.org/a/r912nFRzg0gzmW0J)
