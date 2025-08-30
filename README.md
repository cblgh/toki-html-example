# `toki-html-example`

Example of how to use [`toki`](https://github.com/romshark/toki/) for powering a go-based web application.

For more context, see this issue: https://github.com/romshark/toki/issues/4#issuecomment-3239455700

## Running

Show English translation:

``` 
go run . 
```

Show Swedish translation:

``` 
go run . -lang sv
```


## Updating

```
./update-toki-with-langs.sh
:/update-toki.sh`
go run .
```

## Interesting parts

### `html/index.html`
Note how all the strings in there are basically English and no made up keys.

### `tokibundle/catalog_en.arb`
This contains the strings straight up from `html/index.html` i.e. the translation keys
in `html/index.html` are mapped 1:1 with the values for each English translation key in this
`.arb` file. 

### `tokibundle/catalog_sv_se.arb`

This contains the Swedish translations. The keys (`msg29f53e485d987b3a`, etc) were generated
but I had to fill in the values myself.

### `translations/translations.go`

This is a hacky little program that scans `html/*.html` for invocations of
`.Translator.String`. It records all instances, and then outputs them into
`generated-translations.go`. This is a work-around to get toki to realize that the HTML files
are full of the [TIKs](https://romshark.github.io/tik-cheatsheet/) that power toki—toki
currently only automatically scans `.go` files to find them.

For strings that have parameters, this program uses the following conventions for parameter names to make sure that the
generation proceeds as expected:

* `.Time` for time.Time
* `.NeutralName` contains gender neutral `{name}` (i.e. `tokibundle.GenderNeutral`)
* `.Action` contains arbitrary `{text}

Hacky, but works.
