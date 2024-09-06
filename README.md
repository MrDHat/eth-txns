# eth-txns

EVM blockchain parser that will allow quering transactions for subscribed addresses.

## Usage

### Build the project

```
go build -o eth-txns
```

### Run the project

```
./eth-txns
```

> To run in debug mode, set the `DEBUG` environment variable to `true`. This will print out debug logs to the console.

```
DEBUG=true ./eth-txns
```

### Interacting with the program

Once the program starts up, you can interact with it using the following commands:

#### Subscribe to an address

```
subscribe <address>
```

#### Get transactions for an address

```
get-transactions <address>
```

#### Get the current block number

```
get-current-block
```
