# go-zip-unicode
GoLang library that allows to read Unicode file names in ZIP archives

## Usage
```go
unzipper, err := zip.NewReader(file, fileSize)
for _, f := range unzipper.File {
    correctFileName := zip_unicode_name.ParseUnicodeFileName(f)
}
```
