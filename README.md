# Portfolio

Minimal portfolio website using static HTML/CSS and Go templates. You can clone this repo and customize it to your content and needs.

## Usage

Clone or fork this repo and modify content defined in `blocks.go` file to see your custom content in the portfolio. Run these commands after adding your content:

```sh
# Build binary
make
# Generate static files
make run

# Verify generate files
python3 -m http.server -d public
# Or use your favorite tool to serve static files
```

If you're planning to use Github pages for deployments, please select folder as `public` in Settings section. Also, checkout `make deploy` command in provided Makefile.

#### Credits

I copied design and styling from a github page of https://github.com/tybenz and modified it to work with Go templates.