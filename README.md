# sendx-backend-IIT2020256

# Web Crawler

This is a simple web crawler application that allows you to crawl web pages and view the content. It provides support for concurrent crawling, retrying unavailable pages, and prioritizing paying customers.

## Getting Started

These instructions will help you get the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you start, ensure you have the following installed:

- [Go](https://golang.org/dl/)
- A web browser (for running the client-side of the application)

### Installing

1. Clone the repository to your local machine.

   ```bash
   git clone https://github.com/Ashish7973/sendx-backend-IIT2020256
   cd sendx-backend-IIT2020256
   
2. Start the Go server. Open a terminal and navigate to the project directory, then run:

   ```bash
   go run main.go
   
   
3. The server should now be running on port 5500.
   
4. Open a web browser and enter the following URL to access the web interface:
   
  ```bash
   http://localhost:5500
  ```

# Using the Web Crawler
1. Access the web interface at http://localhost:5500.
2. Enter the URL you want to crawl in the input field.
3. Click the "Crawl" button to initiate the crawling process.
4. The results will be displayed on the new page.

# Advanced Features

## Prioritizing Paying Customers

To prioritize paying customers, add the isPaying query parameter to the URL when making a request. For example:

```bash
http://localhost:5500/crawl?url=https://example.com&isPaying=true
```

# Concurrent Crawling
The web crawler supports concurrent crawling. You can make multiple crawl requests from different browser tabs or clients, and the server will process them concurrently.


