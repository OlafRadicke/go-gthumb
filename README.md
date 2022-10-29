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
go run ./example/example.go ./assets/.comments/Boot-und-Fun_Messe-Berlin.jpg.xml
```

## Status

### Ready implemented

- Realding the files in the .comment drictory :
  - [CommentsRead(path string) (*XMLComment, error)](gthumb.go)
- Writing the files in the .comment drictory:
  - [CommentsWrite(path string, comment *XMLComment) (error)](gthumb.go))

### Todo

- Add interface
- Realding the files bookmarks.xbel
- Realding the files filters.xml
- Realding the files tags.xml
- Writing the files in the .comment drictory
- Writing the files bookmarks.xbel
- Writing the files filters.xml
- Writing the files tags.xml