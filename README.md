# RSS Aggregator

## Overview
The RSS Aggregator is a Go-based application that aggregates RSS feeds, allowing users to fetch, parse, and display updates from multiple sources. This project is a practical implementation of a tutorial available on YouTube, showcasing how to build an RSS aggregator from scratch. It demonstrates core concepts of Go programming, including HTTP handling, concurrency, and database integration.

---

## Features
- **RSS Feed Fetching**: Retrieve and parse RSS feeds from multiple sources.
- **Feed Aggregation**: Combine and display updates from all configured feeds.
- **Concurrency**: Efficiently handle multiple RSS feeds using Go routines.
- **Database Integration**: Store and retrieve feed data using SQL.
- **Web Scraping**: Extract additional information from web pages linked in the RSS feeds.
- **REST API**: Provide endpoints for interacting with the aggregated feeds programmatically.

---

## Prerequisites
- **Go**: Ensure Go is installed on your system. You can download it from [golang.org](https://golang.org/dl/).
- **Database**: A SQL database such as PostgreSQL or MySQL (configurable in the project).

---

## Installation

### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/CTZNpk/rssAggregator.git
    cd rssAggregator
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure the database:
    - Set up a SQL database (e.g., PostgreSQL or MySQL).
    - Update the database configuration in `config/config.go`.

4. Run the application:
    ```bash
    go run main.go
    ```

---


---

## How It Works
- **Feed Fetching**: The app periodically fetches RSS feeds from configured sources and parses the XML to extract relevant data.
- **Aggregation**: Data from all feeds is stored in the database and displayed in a combined format.
- **Concurrency**: Go routines and channels ensure efficient handling of multiple feeds.
- **Web Scraping**: Additional metadata is extracted from linked pages using the scraper module.
- **REST API**: Provides endpoints to access aggregated data programmatically.

