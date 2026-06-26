# ASCII-Art-Web-Stylize

## Description

A continuation of the ASCII-Art-Web project. This version focuses on improving the look and feel of the site — the CSS was pulled out into its own dedicated stylesheet instead of living inline in the HTML, with more styling added throughout. A small piece of JavaScript was also added to clear the input field with a button click instead of manually deleting text. It follows the same principles as the original ascii-art-web project, with the same three banner styles (standard, shadow, thinkertoy).


## Usage

**Clone the repository:**
```bash
git clone <repo-url>
cd ascii-art-web-stylize
```

**Run the server:**
```bash
go run .
```

**Open in your browser:**
```
http://localhost:8080
```

Enter your text, select a banner, and click **Generate Art**. Use the **Clear** button to reset the input field instantly.

## Implementation Details

### Algorithm

The ASCII art generation logic is unchanged from the base project — characters are read from the banner `.txt` files and rendered 8 rows at a time.

### Serving static files

CSS now lives in its own file under a `static/` folder instead of inline in the HTML. To make Go serve it, a file server is mounted on the `/static/` route:

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static", fs))
```

- `http.Dir("./static")` — points the file server at the actual folder on disk where the CSS lives.
- `http.FileServer(...)` — creates a handler that serves any file inside that folder.
- `http.StripPrefix("/static", fs)` — removes `/static` from the incoming URL before looking up the file, so a request to `/static/style.css` correctly maps to `static/style.css` on disk rather than `static/static/style.css`.

The stylesheet is then linked in the HTML like any normal external CSS file:
```html
<link rel="stylesheet" href="/static/style.css" />
```

### Clearing the input field

A small JavaScript function resets the textarea without needing a page reload:

```javascript
function clearInput() {
    document.getElementById("text").value = "";
}
```

```html
<button type="button" onclick="clearInput()">Clear</button>
```

`type="button"` is important here — without it the button would submit the form instead of just running the JS function.

### HTTP Endpoints

| Method | Route | Description |
|---|---|---|
| GET | `/` | Serves the main page |
| POST | `/ascii-art` | Receives text and banner, returns the ASCII art result |
| GET | `/static/*` | Serves static CSS/JS files |

### HTTP Status Codes

| Code | Meaning | When |
|---|---|---|
| 200 | OK | Request handled successfully |
| 400 | Bad Request | Empty text or missing banner |
| 404 | Not Found | Invalid route or banner file not found |
| 500 | Internal Server Error | Templates failed to load at startup |

## What I Learned

This project taught me how to properly separate concerns instead of cramming everything into one HTML file. I learned how Go serves static files — `http.FileServer` combined with `http.StripPrefix` is what lets a URL path like `/static/style.css` map to a real file living in a `static/` folder on disk. Without `StripPrefix`, Go would try to look for `static/static/style.css` and fail, since the URL prefix and the actual folder path aren't automatically the same thing.