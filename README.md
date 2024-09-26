# re-partners
A simple API to calculate the number of packs needed to fulfill customer orders.

## Requirements
- **Go 1.22 or higher**. If you don't have Golang installed, you can download it from [https://go.dev/doc/install](https://go.dev/doc/install)
- **Docker** (optional, for containerized deployment): [https://docs.docker.com/engine/](https://docs.docker.com/engine/)

## Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/ivofreitas/re-partners.git
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Copy local env files**:
    ```bash
    cp internal/config/example.env internal/config/.env
    ```

4. **Run the server**:
    ```bash
    go run cmd/main.go
    ```

5. **Testing with coverage**:
    ```bash
    go test -cover ./...
    ```

## Usage

### UI
Access https://re-partners-nameless-leaf-3542.fly.dev/swagger/index.html to try out it

### Terminal

- **POST /fulfillment/items/calculate-packs**  
  Calculate the number of packs needed for the given number of items.

  Example usage with `curl`:
  ```bash
    curl --location 'https://re-partners-nameless-leaf-3542.fly.dev/fulfillment/items/calculate-packs' \
    --header 'Content-Type: application/json' \
    --data '{
    "total_items": 12001
    }'


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.