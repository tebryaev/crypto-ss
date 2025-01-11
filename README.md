# Seedler

Seedler is a simple Go application that shuffles a seed phrase based on a password. This can be useful for securely storing and sharing seed phrases, as the original order can only be restored with the correct password.

## Features

- **Shuffle Seed Phrase**: The application takes a seed phrase and a password, and shuffles the seed phrase based on the SHA-256 hash of the password.
- **Restore Order**: The original order of the seed phrase can be restored using the same password.
- **Security**: Without the correct password, it is computationally infeasible to restore the original order of the seed phrase.

## Installation

To install Seedler, follow these steps:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/tebryaev/seedler.git
    cd seedler
    ```

2. **Build the application**:
    ```sh
    make build
    ```

3. **Install the binary**:
    ```sh
    sudo make install
    ```

## Usage

To shuffle a seed phrase, run the application and enter the seed phrase and password when prompted:

```sh
seedler
```

Example output:
```
Please enter the password string: 
1 ->  5
2 ->  3
3 ->  8
4 ->  1
...
```

## Security Analysis

The security of the shuffled seed phrase relies on the strength of the password and the SHA-256 hashing algorithm. Here are some key points:

- **Password Strength**: A strong password is crucial. Use a combination of upper and lower case letters, numbers, and special characters.
- **SHA-256 Hashing**: The application uses SHA-256 to generate a hash from the password. This hash is then used to shuffle the seed phrase. SHA-256 is a cryptographic hash function that is resistant to pre-image and collision attacks.
- **Probability of Restoration**: Without the correct password, the probability of correctly guessing the original order of a 24-word seed phrase is \(24!\) (24 factorial), which is approximately \(6.204 \times 10^{23}\). This makes it computationally infeasible for an attacker to restore the original order without the password.
