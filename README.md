go-gthumb
---------

Small library for manipulation of meta and config files of program
[gThumb](https://wiki.gnome.org/Apps/Gthumb).

## Installation

```
go get github.com/OlafRadicke/
```

## Basic usage



## Test and example

Run test with:

```bash
go run example/example_comment_files.go
./scrpts/cean_up_test.sh
go run example/example_tag_files.go
./scrpts/cean_up_test.sh
```

## Status

### Ready implemented

- Realding the files in the .comment drictory
- Writing the files in the .comment drictory
- Add and remove methodes for from the list of categories
- Get methode for version of comment file format
- Get and set method for Node value.
- Get and set method for Place value.
- Get and set method for Caption value.
- Get and set method for Rating value.


### Todo

- Add interface
- Realding the files bookmarks.xbel
- Realding the files filters.xml
- Realding the files tags.xml
- Writing the files in the .comment drictory
- Writing the files bookmarks.xbel
- Writing the files filters.xml
- Writing the files tags.xml

Grüße