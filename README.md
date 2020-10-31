# Portfolio

Minimal portfolio website using static HTML/CSS and Go templates. You can clone this repo and customize it to your content and needs.

See my portfolio here - [jan25.github.io](http://jan25.github.io/).

## Usage

Clone or fork this repo and modify content defined in [`content.go`](content.go) file to see your custom content in the portfolio. Run these commands after adding your content:

```sh
# Generate static files
make build-and-generate
# OR simply
make

# Verify generated files
python3 -m http.server -d public
# Or use your favorite tool to serve static files
```

If you're planning to use Github pages for deployments, checkout `deploy` command provided in [Makefile](Makefile).

#### Credits

I copied design and styling from a github page of [@tybenz](https://github.com/tybenz) and modified it to work with Go templates.