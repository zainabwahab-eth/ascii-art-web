# ASCII-Art-Web-Export-File

## Description

A continuation of the ASCII-Art-Web project. This version adds the ability to export the generated ASCII art as a downloadable `.txt` file, instead of only viewing it in the browser.

## Usage

**Clone the repository:**
```bash
git clone <repo-url>
cd ascii-art-web-export
```

**Run the server:**
```bash
go run .
```

**Open in your browser:**
```
http://localhost:8080
```

Enter your text, select a banner, and click **Generate**. Once the result is on the page, click **Download .txt** to export it.

## Implementation Details

### Algorithm

The ASCII art generation logic is unchanged from the base project — characters are read from the banner `.txt` files and rendered 8 rows at a time.

Unlike a naive implementation where clicking download would re-run the generation logic a second time, this version computes the result once. When the user clicks **Generate**, the result is rendered onto the page and also stored in a hidden input field:

```html
<textarea name="txt-result" id="txt-result" class="hidden">{{.Result}}</textarea>

```

When the user then clicks **Download**, that hidden field is submitted to `/download`, which simply reads the already-computed result back out and sends it as a file — no need to call the ASCII art generator twice for the same input.

```go
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    result := r.FormValue("txt-result")

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(result))
}
```

### The HTTP headers

| Header | Purpose |
|---|---|
| `Content-Type` | Tells the browser what kind of content is being sent (`text/plain`) |
| `Content-Disposition` | Tells the browser to download the response as a file instead of rendering it, and sets the filename |
| `Content-Length` | Tells the browser the exact size of the file in bytes |

### HTTP Endpoints

| Method | Route | Description |
|---|---|---|
| GET | `/` | Serves the main page |
| POST | `/ascii-art` | Generates the ASCII art and renders it on the page |
| POST | `/download` | Reads the already-generated result from a hidden field and returns it as a downloadable `.txt` file |

### HTTP Status Codes

| Code | Meaning | When |
|---|---|---|
| 200 | OK | Request handled successfully |
| 400 | Bad Request | Empty text or missing banner |
| 404 | Not Found | Invalid route or banner file not found |
| 500 | Internal Server Error | Templates failed to load at startup |

## What I Learned

This project shifted my understanding of HTTP responses — I learned that what makes a browser treat a response as a downloadable file rather than a page to render is entirely about the **headers**, not the content itself. The same plain text bytes get written to `w` either way — it's `Content-Disposition: attachment` that tells the browser "save this" instead of "display this."

I also caught an inefficiency in my first version — my original download handler was re-running the entire ASCII art generation a second time just to produce a file, even though the result had already been computed and shown on the page seconds earlier. I fixed this by storing the result in a hidden form field after generating it, so the download button just reads that value back out instead of doing the same work twice. It was a good reminder to think about where data already exists before recomputing it.
