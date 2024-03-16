# Data Migration Tool

This is a simple open-source tool written in Go for migrating data from a CSV file to a PostgreSQL database.

## Requirements

- Connect to a remote PostgreSQL database.
- Read data from a CSV file.
- Insert each row of the CSV file into a PostgreSQL table.
- Allow users to customize column mappings between the CSV file and the PostgreSQL table.
- Ensure fast execution.
- Implement a duplicate check to prevent inserting duplicate entries.

## Version 1 Limitations

- Works only with PostgreSQL.
- Compatible with all operating systems.
- Assumes an existing table in PostgreSQL; does not attempt to create its own.

## Getting Started

### Prerequisites

- Go installed on your system.
- Access to a PostgreSQL database.
- Basic understanding of PostgreSQL and CSV file structure.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/data-migration-tool.git
    cd data-migration-tool
    ```

2. Install dependencies:

    ```bash
    go get github.com/lib/pq
    ```

### Usage

1. Update the PostgreSQL connection details in `main.go`:

    ```go
    const (
        host     = "your_postgres_host"
        port     = 5432
        user     = "your_postgres_user"
        password = "your_postgres_password"
        dbname   = "your_postgres_db"
    )
    ```

2. Place your CSV file (`data.csv`) in the project directory.

3. Customize the code in `main.go` to map CSV columns to PostgreSQL table columns.

4. Run the tool:

    ```bash
    go run main.go
    ```

### Customization

- Modify the INSERT query in `main.go` to match your PostgreSQL table schema.
- Implement additional features such as user-defined column mappings, error handling, etc., as per your requirements.

### Contributing

Contributions are welcome! Feel free to open issues and pull requests.

### License

This project is licensed under the [MIT License](LICENSE).
