# Puzzler

[![License Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

Puzzler is a layout service that puzzles together the pieces of a layout from different services.

Modern day websites are a creation by multiple teams. However, many of these modern websites are still build as monolith. [Ruby on Rails](https://rubyonrails.org/) is one of the frameworks that weaves together one language and one framework to build them. But, there are many more languages than Ruby and more frameworks. Microserivces are build on the idea that you have simpler, more specific components.

## Fragement(s)

A `fragment` is polymorphic. On the server it is parsed and evaluate to a `puzzle`. In the browser it is a web component that can receives streamed data.

### Server

* `src` The source to fetch for replacement in the DOM
* `method` can be of `GET` (default) or `POST`.
* `id` is an optional unique identifier (optional)
* `timeout` timeout of a fragement to receive in milliseconds (default is `300`)
* `deferred` is deferring the fetch to the browser
* `fallback` is deferring the fetch to the browser if failed (default)

### Frontend

This is when the fragment is deferred to the browser

* `src` The source to fetch for replacement in the DOM

## Example

```html
<html>
<head>
    <script type="fragment" src="http://localhost:3000/assets"></script>
</head>
<body>
    <h1>Example</h1>
    <fragment src="http://localhost:3000/fragment1.html"></fragment>
</body>
</html>

```

## License

[Apache 2.0](/LICENSE)
