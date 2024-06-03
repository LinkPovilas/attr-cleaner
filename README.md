# attr-cleaner

A utility for parsing HTML files, stripping all HTML attributes except `data-*`, and saving the modified file.

```bash
# Where * is the filename
go run main.go *.html ./output/*.html
```
