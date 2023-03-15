# Forum

The Forum project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

## Usage/Examples

In this example _forum:1.0_ is the name of Docker image and _forum-ap_ is the name of Docker container
Build Docker image based on Dockerfile with Makefile:

```bash
$ make build
```

Create Docker container:

```bash
$ make run
```

server started at http://localhost:8081/

Have fun!
