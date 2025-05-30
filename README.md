# Go Svelte SPA

This repo contains example code of how to integrate a [Svelte](https://svelte.dev/) single-page application (SPA), built with [SvelteKit](https://svelte.dev/docs/kit/introduction), with a [Golang](https://go.dev/) backend.

In development mode, the Golang backend will proxy requests for the frontend assets to the Vite server (running at http://localhost:5173).

In production mode, the Golang backend will serve the assets bundled from the `ui/dist` directory.

The mode is determined by the `GO_ENV` environment variable. If it contains "prod" then it will run in production mode, otherwise it will run in development mode.

## Motivation

This setup lets you create a single, easy-to-deploy binary that contains both your API and frontend code without sacrificing on the developer experience. It is inspired by [PocketBase](https://pocketbase.io/).

## Requirements

- Go 1.22+
- pnpm (optional, you can replace `pnpm` with `npm` or `yarn` in the Makefile)

## Try it out

Development mode:

```
make watch
```

Production mode:

```bash
make build
GO_ENV=production ./main
```

The application will be accessible at http://localhost:8080

## Development

The following commands use the `examples/stdlib/main.go` file for the backend.

Run build make command with tests

```bash
make all
```

Build the entire application

```bash
make build
```

Build the UI which will output the assets to `ui/dist`

```bash
make build-ui
```

Build the server which will output a `main` executable binary

```bash
make build-server
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## Manually setting up the SvelteKit UI

Follow instructions from the official docs for initializing a new app. At the
time of writing this:

```
npx sv create ui
cd ui
npm install
```

Setup the static adapter:

- Install: `npm install @sveltejs/adapter-static`
- Setup the adapter in `svelte.config.js`:

```js
// ...
kit: {
  adapter: adapter({
    pages: "dist",
    assets: "dist",
    fallback: "index.html",
    strict: true,
  });
}
```

Add a `layout.js` (or layout.ts) file with the following exports to enable SPA routing:

```js
export const ssr = false;
export const prerender = false;
```
