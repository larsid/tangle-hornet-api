# Tangle Hornet API
API written in Go to use Tangle Hornet Network.

## How to use

### Requirements
> This library was mainly tested with Go version 1.20.x

We recommend you update Go [to the latest stable version to use the library](https://go.dev/).

### Downloading the API

1. You can get the latest version of the API by [clicking here](https://github.com/AllanCapistrano/tangle-hornet-api/releases/download/v1.5.0/tangle-hornet-api);
2. Give the permission to execute the API:
   ```powershell
   sudo chmod +x tangle-hornet-api
   ```
3. If you want to change some parameters, create a file named `tangle-hornet.conf` in the `/etc/` directory;
   1. Check the [tangle-hornet.conf](https://github.com/AllanCapistrano/tangle-hornet-api/blob/main/config/tangle-hornet.conf) to see all the parameters you can override.
4. Run the API:
   ```powershell
   ./tangle-hornet-api
   ```
   
### Build the project on your own

1. Clone this repository;
2. Download all the dependencies:
   ```powershell
   go mod tidy
   ```
3. Execute the project:
   ```powershell
   go run main.go
   ```

## Endpoints

| Method | Enpoint | Description |
| ------ | ------- | ----------- | 
| GET | `/nodeInfo` | Shows information about Tangle Hornet Network. | 
| GET | `/nodeInfo/all` | Shows all information about Tangle Hornet Network. | 
| GET | `/message/messageId/{messageId}` | Get a message by given message ID. |
| GET | `/message/{index}` | Get all messages using a specific index. |
| GET | `/message/{index}/{maxMessages}` | Get a maximum number of messages from the last hour, using a specific index. |
| POST | `/message` | Create and submit a new message. |

You also can use the files in the [`apiClientFiles`](./apiClientFiles/) directory to import all routes into your API Client.

### POST Body Examples

#### Routes

- `/message`:
  ```json
  {
		"index": "LB_STATUS",
		"data": {"available": true, "avgLoad": 3, "createdAt": 1695652263921, "group": "group4", "lastLoad": 4, "publishedAt": 1695652267529, "source": "source4", "type": "LB_STATUS"}
	}
  ```

## License
The MIT license can be found [here](./LICENSE).
