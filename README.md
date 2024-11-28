# go-storage-handler
### Description:
Go-Storage-Handler is a command-line interface (CLI) application built to bridge PostgreSQL and MongoDB, enabling efficient data storage and retrieval in binary format. This tool is designed for developers and data engineers who require a streamlined, high-performance solution for handling data across relational and NoSQL databases. The CLI offers essential functionality for data migration, synchronization, and storage management in distributed database environments.

[Watch "Go Storage Handler CLI Application" on YouTube](https://www.youtube.com/watch?v=giLDNG9wu9Q)

![image](https://github.com/user-attachments/assets/a888f44f-9968-4964-b7df-e75cb588653f)

## CLI Commands
### Key Commands:

* upload: Uploads a specified file to the target storage location in PostgreSQL or MongoDB.
* update: Modifies an existing file in storage with new data.
* download: Downloads a file from storage, supporting binary data retrieval.
* delete: Removes a specified file from storage, freeing space and ensuring storage efficiency.
* help, h: Displays a list of available commands or detailed help for a specific command.

### Global Options:
* --configPath : Defines the path to the config.json configuration file (default: "config.json").
* --help, -h : Provides help documentation.
* --version, -v : Displays the current version of Go-Storage-Handler.

### Sub Flags :
* --file (value) : Specifies the file name or file path to be used with upload, update, download, and delete commands.

## Make Commands
* make build : Compiles the application for the local OS and architecture.
* make clean : Cleans up build artifacts by removing the build directory.
* make test : Executes all tests.
* make build-linux : Builds the application for Linux.
* make build-darwin : Builds the application for macOS.
* make build-windows : Builds the application for Windows.
* make build-all : Compiles the application for all major platforms (Windows, macOS, Linux).
* make help : Shows a summary of available make commands.

## Features:

#### Cross-Database Connectivity:

* Seamlessly connects to multiple PostgreSQL and MongoDB databases.
* Supports simultaneous connections to different database instances, enabling unified data operations across databases.
Binary Data Storage and Retrieval:
* Transfers data in binary format, optimizing storage and minimizing data transformation overhead.
* Uses PostgreSQL’s BLOB/CLOB types and MongoDB’s GridFS or binary storage, ensuring compatibility and performance.
#### Intuitive CLI Commands:

* Provides user-friendly commands for data transfer, configuration, and management.
* Enables filtering of data and fine-tuning of parameters for more controlled data synchronization.
  
#### Data Synchronization and Migration Support:
* Offers capabilities for full data migration, incremental updates, or periodic synchronization.
* Ideal for setting up regular data transfer jobs, data backup, and distributed database consistency.
  
#### Robust Logging and Error Handling:
* Includes detailed logging for tracking operations and troubleshooting errors.
* Verbose and debug modes allow visibility into the data transfer process, enhancing transparency and control.

### Exmple Commands:
* ./goStorageHandler-windows.exe --configPath=config.json upload --file=/path/x.png
* ./goStorageHandler-windows.exe --configPath=config.json download --file=x.png
* ./goStorageHandler-windows.exe --configPath=config.json update --file=/path/x.png
* ./goStorageHandler-windows.exe --configPath=config.json delete --file=x.png

##
Go-Storage-Handler is an essential tool for database professionals needing an efficient and secure solution for data management in multi-database environments. It simplifies complex tasks and offers robust support for cross-database data storage, retrieval, and management across PostgreSQL and MongoDB.
