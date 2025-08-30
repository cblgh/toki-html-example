when running toki-prep.sh, i had this error 

[ERROR] analyzing sources: errors in package "tokibundle": [/home/cblgh/code/toki-golang-html-i18ning/tokibundle/bundle_gen.go:14:2: could not import golang.org/x/text/language (invalid package name: "")]

it went away after i did:

```
go get github.com/go-playground/locales
go get golang.org/x/text/language
```

## questions 

* how do i remove a translation language that was added by mistake?
* how should i best match entries in a new translation's `.arb` with the actual words that need
  to be translated, should i just copy the keys from the bottom of `bundle_gen.go` and compare
  them to the empty ICUs in the `.arb` file for the new translation?
* will you add scanning of `.html` files for calls like {{ .tokiTranslator.String "TIK here" }}?
