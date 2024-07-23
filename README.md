# GoGet - Go Package Manager

GoGet is a simple package manager for Go written in Go. It allows you to install, update, and remove Go packages from the command line.

## Features

- **Install**: Install Go packages from the command line.
- **Update**: Update installed Go packages.
- **Remove**: (Placeholder) Remove Go packages.

## Installation

To build and use GoPM, follow these steps:

1. **Clone the repository:**

    ```sh
    git clone https://github.com/xanmoy/goget.git
    cd goget
    ```

2. **Build the project:**

    ```sh
    go build -o goget ./cmd/gopm
    ```

3. **Install the binary (optional):**

    You can move the `goget` binary to a directory in your `PATH` for easier access:

    ```sh
    sudo mv goget /usr/local/bin/
    ```

## Usage

Once installed, you can use GoGet from the command line with the following commands:

- **Install a package:**

    ```sh
    goget install <package>
    ```

    Example:

    ```sh
    goget install github.com/sirupsen/logrus
    ```

- **Update a package:**

    ```sh
    goget update <package>
    ```

    Example:

    ```sh
    goget update github.com/sirupsen/logrus
    ```

- **Remove a package:**

    ```sh
    goget remove <package>
    ```

    Example:

    ```sh
    goget remove github.com/sirupsen/logrus
    ```

    (Note: The remove functionality is not yet implemented.)

## Contributing

Contributions are welcome! If you have suggestions or improvements, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or issues, please open an issue on GitHub or contact [contact@xanmoy.me](mailto:contact@xanmoy.me).

