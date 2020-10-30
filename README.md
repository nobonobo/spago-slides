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

<li class="fragment">Item1</li>
<li class="fragment">Item2</li>

====

# Last Page
```

- `====`: page separator.
- class="fragment": not visible at first.

## Step3. execute server

```shell
spago-slides content.md
```

default listen port is ":8080"

## Step4. open on your browser

open htttp://localhost:8080

## Step5. start your presantation!

key map

- Arrow Right(▶️buttton click):
  If the slide has fragments, they will be displayed, otherwise the slide move to next page.
- Arrow Left(◀️buttton click):
  the slide move to prev page.
- `R` key:
  reload content.md

## Step extra.

If you use custom.css, you can use `-css` option.

```shell
spago-slides -css custoom.css content.md
```

## Enjoy!
