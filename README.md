# search-engine-BE




This is the backend service for a full-stack search engine built using Go and Fiber.  
It loads data from a Parquet file, stores it in memory, and provides APIs to upload new files and perform search operations.






# Architecture Overview

- Server Framework - [Fiber]- a fast and Express-like web framework for Go.
- File Handling - [Parquet-Go]-library to parse `.parquet` files.
- In-memory Storage - Custom slice (`[]Record`) loaded with parsed data.
- Search Mechanism -  Simple in-memory linear search with basic text matching on fields like `Message`, ect....





 # APIs

  - `POST /upload` — Uploads a new Parquet file and loads its contents.
  - `GET /search?query=...` — Searches records containing the query text.





# Design Choices

The backend was designed with a focus on simplicity, speed, and adherence to assignment constraints (no external database or search engine allowed). Here’s a breakdown of the key architectural decisions:

   # Web Framework:
    We use Fiber — a fast, minimalist web framework for Go that offers an Express.js-like API. Fiber was chosen for its speed, scalability, and ease of building RESTful APIs quickly.

   # File Storage Strategy:
    Instead of persisting data into an external database, all parsed Parquet records are loaded into an in-memory slice ([]Record). This approach ensures extremely fast read/search operations for small to medium-sized datasets without the complexity of database management.

   # File Format:
    Parquet files are used as the data source. Parquet is a highly efficient columnar storage format, making it ideal for structured logging or event data.

   # Search Logic:
    A basic substring matching approach is used to scan text fields like Message, Tag, and Sender. Although simple, this method is sufficient and performant given the dataset size and memory-based architecture.

   # Concurrency Handling:
    The Fiber framework, combined with Go's native support for goroutines, allows the backend to handle multiple concurrent upload and search requests efficiently without blocking the server.

   # No External Database:
    All data is kept entirely in memory. This was a conscious design decision to align with the assignment's rule of no third-party database usage (e.g., Postgres, MongoDB, etc.), ensuring the system remains lightweight and self-contained.





# Build and Run the Application

1. Go 1.23.3 windows/amd64 installed. 
2. A .parquet file available for testing.
3. Install Go Dependencies -- go mod tidy
4. Run the Backend Server -- go run main.go
5. If successful, you will see output like--Server running at http://localhost:4000
6. Test the API Endpoints --  http://localhost:4000/upload
7. Search for a keyword -- http://localhost:8080/search?query=your_keyword


# Benchmarking Approach

1. Load Time:
   The application measures the time taken to parse and load the Parquet file into memory upon upload.

2. Search Time:
   Each search API call (/search?query=...) logs the time taken to perform the search operation using Go's time package.
   e.g
   start := time.Now()
   elapsed := time.Since(start)



# Performance Tuning Strategies

1. In-Memory Storage -- 	Using an in-memory slice ([]Record) avoids database overhead, ensuring fast reads.
2. Debounced Search (Frontend) -- 	Frontend sends search API requests with a 0.5-second debounce to avoid overloading the backend.
3. Concurrent Parquet Reading -- 	Parquet reader uses 4 parallel goroutines (NewParquetReader(fr, new(search.Record), 4)) to speed up file       parsing.
4. Efficient Field Matching -- Search is performed only on selected important fields to minimize processing per record.


# goals or notable enhancements implementetion.

1. Dynamic File Upload Support -- The backend accepts new Parquet files at runtime without needing a server restart.
                               -- Users can upload a new dataset via the /upload API, and the in-memory data gets refreshed immediately.

2. Search Performance Measurement -- Every search API call logs the exact time taken to perform the search.
                                  -- Search time is also included in the API response (searchTime field) to allow easy frontend display and        performance tracking.
            
3. Debounced Search in Frontend -- Introduced a 2-second debounce mechanism on the frontend search field.
                                -- This reduces unnecessary backend load and improves overall user experience during fast typing.

4. Error Handling and Validation -- Graceful error responses if no file is uploaded or if search queries are missing.



