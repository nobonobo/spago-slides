# spago-slides

Web-base Presantation CLI tool for Gopher.

## Step1. install

```shell
go get github.com/nobonobo/spago-slides
```

## Step2. write markdown

content.md

```markdown
# Title

====

# Next Page
```

`====` is 　 page separator.

## Step3. execute server

```shell
spago-slides content.md
```

default listen port is ":8080"

## Step4. open on your browser

open htttp://localhost:8080

## Step5. start your presantation!

key map

- Arrow Right(▶️buttton click): appear fragments or move to next page.
- Arrow Left(◀️buttton click): move to prev page.

## Step extra.

If you use custom.css, you can use `-css` option.

```shell
spago-slides -css custoom.css content.md
```

## Enjoy!
