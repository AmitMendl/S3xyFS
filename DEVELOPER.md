# Internal docs

## code structure

The S3 interpretation goes through the following pipeline

server package (gets requests) => S3 package (transforms request) => FS package (processes requests)

after the request has been processed (read, write, etc...). the same pipeline repeats itself

FS package (creates response) => S3 package (transforms response) => server package (sends response)
