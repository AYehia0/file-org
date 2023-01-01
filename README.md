# File-Org
The aim of this program is to learn go while also making a useful working CLI application that manages files the proper way.

# How does it work ?
It's a very simple application, that takes 2 input `args`.
```
file-org /path/to/scanning/location .
```
- Scanning Location : the location where the messy files exist (ex, `"~/Downloads"`)
- Target Location : the location to organize the files in.

It works by scanning for all the `files` in that scanning location then returns a `map` of all the files with extension types as keys.
If the extension name is in the list of wanted extensions, it's moved!

# Installation

Make sure you have go installed.
```
go install github.com/AYehia/file-org@latest
```

# TODO
- [ ] Add Tests.
- [ ] Use structs.
- [ ] Remove some dependencies.
- [x] Make use of the really nice `vim-go` docs
- [ ] Log file with added files.

# Resources

- [The Little Go Book](https://www.openmymind.net/The-Little-Go-Book/)
- [Rob's awesome-go](https://github.com/rwxrob/awesome-go)
