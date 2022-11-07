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

- files in the .comment directory
  - Read the files in the .comment directory
  - Write the files in the .comment directory
  - Add and remove methodes for from the list of categories
  - Get method for version of comment file format
  - Get and set method for comment node value.
  - Get and set method for comment place value.
  - Get and set method for comment time value.
  - Get and set method for comment caption value.
  - Get and set method for read comment rating value.
- ~/.config/gthumb/tags.xml
  - Read the file with the tags configuration.
  - Write the file with the tags configuration.
  - get, add und remove tags from the tags configuration file.

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