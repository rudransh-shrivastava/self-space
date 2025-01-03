# Self Space

Self Space is a self hosted alternative to AWS S3


# Example Usage
You can also see all the commands [here](#available-cli-commands)

### Clone the Repository:
```bash
git clone https://github.com/rudransh-shrivastava/self-space
```

### Change Directory
```bash
cd self-space
```

### Install Dependencies
```bash
go mod tidy
```

### Build

if you have `make` installed you can:
```bash
make build
```
if you don't then you can:
```bash
go build -o bin/self-space
```

### Setup Environment Variables

create a `.env` file.
```bash
touch .env
```

Available Environment Variables:
```
PUBLIC_HOST="localhost"
PORT="8080"
BUCKET_PATH="buckets/"
BUFFER_SIZE=1024
BUCKET_NAME_MAX_LENGTH=100
```
You can copy these default environment variables to your `.env` file. If you don't create a `.env` file, these defaults will be used by the server.

### Change Directory to Build
```bash
cd bin
```

### Generate an API Key
```bash
./self-space apikey generate
```
Example output:
```
generated api key: 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ=
```
Save this API Key, as you will need it to call the APIs and to attach a bucket.

### Create a Bucket
```bash
./self-space bucket create myBucket
```
Example output:
```
created bucket with name: myBucket
```

### Attach API Key to the Bucket

replace `3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ=` with the API Key you saved earlier.
```bash
./self-space apikey attach 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= myBucket READ
```
Example output:
```
attached bucket myBucket to api key 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= with permission READ
```
Since, we also want to Download and Delete our file in this example, we will give it two more permissions, `WRITE` and `DELETE`
```bash
./self-space apikey attach 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= myBucket WRITE
```
```bash
./self-space apikey attach 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= myBucket DELETE
```
Example outputs:
```
attached bucket myBucket to api key 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= with permission WRITE
attached bucket myBucket to api key 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ= with permission DELETE
```

### Start the Server
```bash
./self-space
```
Example output:
```
listening on localhost:8080
```

### Upload a File

Let's make a `PUT` request using `curl` to upload a `file.txt` to our bucket `myBucket`
first create a `file.txt` and write something in the file.

then, we can hit our API, and upload the file to out bucket.
```bash
 curl -X PUT -T file.txt http://localhost:8080/bucket/myBucket/directory/file.txt \
             -H "X-API-Key: 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ="
```
You will see a `buckets/myBucket` directory being created inside the `bin` directory
Our file is uploaded to `bin/buckets/myBucket/file.txt`

### Download a File

Let's download the file we just uploaded earlier.
```bash
curl -X GET http://localhost:8080/bucket/myBucket/file.txt \
     -H "X-API-Key: 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ=" -o downloaded.txt
```
Example output:
```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    91  100    91    0     0   1708      0 --:--:-- --:--:-- --:--:--  1716
```
this will download our file as `downloaded.txt`

### Delete a File

We can also delete the same file.
```bash
curl -X DELETE http://localhost:8080/bucket/myBucket/file.txt \
     -H "X-API-Key: 3tVHFOjY6mpfEtb7si4y_EudHOH322WvBSQ1qQOnQSQ="
```
Example output:
```
File deleted successfully: buckets/myBucket//file.txt
```

# Available CLI Commands

## Run Server
Runs the server
```bash
self-space
```

## API Key Commands
The `apikey` command is used to manage API keys.

### Generate a New API Key

```bash
self-space apikey generate
```
**Description**: Generates a new API key.

### Attach a Bucket to an API Key
```bash
self-space apikey attach <api-key> <bucket-name> <permission>
```
**Description**: Attach a bucket to an API key with a given permission.
**Usage**:

-   `<api-key>`: The API key to attach the bucket to.
-   `<bucket-name>`: The name of the bucket to attach.
-   `<permission>`: The permission to assign (`READ`, `WRITE`, or `DELETE`).

### Detach a Bucket from an API Key
```bash
self-space apikey detach <api-key> <bucket-name> <permission>
```
**Description**: Detach a bucket from an API key completely or for a specific permission.
**Usage**:

-   `<api-key>`: The API key to detach the bucket from.
-   `<bucket-name>`: The name of the bucket to detach.
-   `<permission>`: The permission to remove (`READ`, `WRITE`, or `DELETE`).

### Delete an API Key
```bash
self-space apikey delete <api-key>
```
**Description**: Deletes the specified API key.
**Usage**:

-   `<api-key>`: The API key to delete.

### List All API Keys
```bash
self-space apikey list
```
**Description**: Lists all API keys.

## Bucket Commands

The `bucket` command is used to manage buckets.

### Create a Bucket
```bash
self-space bucket create <bucket-name>
```
**Description**: Creates a new bucket.
**Usage**:

-   `<bucket-name>`: The name of the bucket to create.

### Delete a Bucket
```bash
self-space bucket delete <bucket-name>
```
**Description**: Deletes the specified bucket.
**Usage**:

-   `<bucket-name>`: The name of the bucket to delete.

### List All Buckets
```bash
self-space bucket list
```
**Description**: Lists all available buckets.
