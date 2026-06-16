# ASCII-Art-Web

## Description

A web-based GUI version of the ASCII-Art project built with Go's standard library. It allows you to convert any text into ASCII art directly from the browser using three different banner styles. The result is displayed on the same page after submission.

## Usage

**Clone the repository:**

```bash
git clone <repo-url>
cd ascii-art-web
```

**Run the server:**

```bash
go run .
```

**Open in your browser:**

```
http://localhost:8080
```

Enter your text, select a banner style, and click **Generate Art**.

## Implementation Details

### Algorithm

The server reads the banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) from the `banners/` directory. Each file contains all printable ASCII characters (32–126) represented as 8-line tall art blocks separated by a blank line.

To render the ASCII art:

1. The input text is split by newline into separate lines
2. For each line, the program iterates through every character
3. For each character it calculates its position in the banner file using the formula: `index = (charCode - 32) * 9 + 1`
4. It reads the 8 rows for that character and builds the output row by row
5. The final result is passed to the HTML template and displayed in a `<pre>` block

### HTTP Endpoints

| Method | Route        | Description                                            |
| ------ | ------------ | ------------------------------------------------------ |
| GET    | `/`          | Serves the main page                                   |
| POST   | `/ascii-art` | Receives text and banner, returns the ASCII art result |

### HTTP Status Codes

| Code | Meaning               | When                                   |
| ---- | --------------------- | -------------------------------------- |
| 200  | OK                    | Request handled successfully           |
| 400  | Bad Request           | Empty text or missing banner           |
| 404  | Not Found             | Invalid route or banner file not found |
| 500  | Internal Server Error | Templates failed to load at startup    |

### Project Structure
