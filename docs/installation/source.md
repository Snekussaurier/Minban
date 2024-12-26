# Source Installation

To build and run the MinBan backend from source, follow these steps:

## 1. Clone the Repository

First, clone the MinBan backend repository:

```bash
git clone https://github.com/Snekussaurier/minban-backend.git
cd minban-backend
```

## 2. Install Dependencies

!!! info
    Make sure you have Go 1.23+ and SQLite installed on your system.

#### For **Go** installation:
Follow the instructions at golang.org.

#### For **SQLite** installation:
Refer to your system's package manager:

=== "Ubuntu/Debian"

    ``` bash
    sudo apt install sqlite3
    ```

=== "MacOS"

    ``` bash
    brew install sqlite
    ```

Additionally, install gcc and musl-dev for building Go applications

## 3. Build the Application

Navigate to the src folder:

``` bash
cd src
```

Download the Go module dependencies:

``` bash
go mod download
```

Build the application binary:

```bash
go build -o main .
```

This will create an executable file named main in the current directory.

## 4. Environment Variables
The application needs the following environment variables to be configured. 

- **DATABASE_PATH:** Path to the SQLite database file.
- **USER_NAME:** Admin username.
- **USER_PASSWORD:** Admin password.
- **JWT_SECRET_KEY:** Secret key for JWT authentication.


## 5. Run the Server

Once the binary is built, navigate back to the project root and run the server:

``` bash
./main
```

!!! info
    The server will start on port 9916 by default. You can access the API at http://localhost:9916.

