# Flint
Language-Agnostic Static Sites

`flint` is a static site generator that works regardless of the way you build your application. Run you app on `localhost`, setup your `flint.json` and run `flint spark` to generate your static assets. 

## Installation
To install, clone the repo:
```bash
git clone https://github.com/phillip-england/flint
```

Then, with go `1.23.3` or later:
```bash
go build -o flint main.go
```

Then, you can move `flint` somewhere on your `PATH` to use on your system.

## Config
`flint` require a `flint.json` to gain a bit of context on how to build your assets.

Here is an example `flint.json`:
```json
{
    "host": "http://localhost:8080",
    "static": "./static",
    "favicon": "./favicon.ico",
    "out": "./out",
    "target": "https://phillip-england.github.io/www.stacijs.com",
    "routes": [
        "/",
        "/docs/signals",
        "/docs/events",
        "/docs/observers",
        "/docs/installation"
    ]
}
```

Let's break down each item.

### Host
The `host` is simply where your application is running. Right now, `flint` focuses on building local applications, but it could be extended to generate static sites from MPA running on the web.

### Static
`static` tells `flint` where to find static assets associated with the running application. These assets will be bundled and minified during the generation process.

### Favicon
`favicon` tells `flint` where to find the `favicon.ico` for the application.

### Out
`out` lets `flint` know where to place the static assets after generation.

### Target
`target` tells `flint` where the application will be deployed. This enables flint to crawl all of the relative `link=` and `src=` attributes on elements and change them from relative paths to absolute paths. This enables all links are properly transformed and work as expected in the target environment.

### Routes
`routes` tells `flint` which endpoints to hit during static site generation.

## Running
To generate static assets, run `flint spark`. From there, you can take the `out` directory and plop it wherever you deploy your static sites. Enjoy.