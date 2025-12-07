![Go Version](https://img.shields.io/badge/go-1.25+-blue)
![OS](https://img.shields.io/badge/OS-Windows,Linux,MacOS-orange)
![License](https://img.shields.io/badge/license-MIT-green)

# Table of Content

- [About](#about)
- [Features](#features)
- [Performance](#performance)
- [Getting Started](#getting-started)
- [Architecture](#architecture)
  - [Core System Components](#core-system-components)
  - [Execution Flow](#execution-flow)
- [Examples](#examples)

# About

Gopher Cache is a small in-memory database. It was built with the intention to mock a Redis database, so that it behaves in a similar way as the real one from the user/client perspective. The source code of the project is fully written in Go. The initial goal of this project was to build a mini copy of the Redis database in order to increase understanding of both the Redis internals and the Go programming language. However, everybody is welcome to check the code, propose improvements, create their own forks/versions, and use the source code for their own purposes.

> Note: This project is not intended to be a drop-in replacement for Redis.

# Features

- **RESP Command Parser:** Recursive-descent parser capable of parsing RESP arrays of bulk strings.
- **TCP Server:** Listens for incoming client connections and queues incoming commands for execution.
- **FIFO Command Execution:** Commands are executed in first-in-first-out order as they arrive to the server.
- **Redis commands support:**
  | Command | Redis Docs | Supported | Notes |
  | ------------- |:-------------| :-----:| ---- |
  | SET | [https://redis.io/docs/latest/commands/set](https://redis.io/docs/latest/commands/set/) | ⚠️ | (supported without options) |
  | GET | [https://redis.io/docs/latest/commands/get](https://redis.io/docs/latest/commands/get/) | ✅ | |
  | HSET | [https://redis.io/docs/latest/commands/hset](https://redis.io/docs/latest/commands/hset/) | ✅ | |
  | HGET | [https://redis.io/docs/latest/commands/hget](https://redis.io/docs/latest/commands/hget/) | ✅ | |
  | HMGET | [https://redis.io/docs/latest/commands/hmget](https://redis.io/docs/latest/commands/hmget/) | ✅ | |
  | LPUSH | [https://redis.io/docs/latest/commands/lpush](https://redis.io/docs/latest/commands/lpush/) | ✅ | |
  | RPUSH | [https://redis.io/docs/latest/commands/rpush](https://redis.io/docs/latest/commands/rpush/) | ✅ | |
  | LPOP | [https://redis.io/docs/latest/commands/lpop](https://redis.io/docs/latest/commands/lpop/) | ✅ | |
  | RPOP | [https://redis.io/docs/latest/commands/rpop](https://redis.io/docs/latest/commands/rpop/) | ✅ | |
  | LLEN | [https://redis.io/docs/latest/commands/llen](https://redis.io/docs/latest/commands/llen/) | ✅ | |
  | SADD | [https://redis.io/docs/latest/commands/sadd](https://redis.io/docs/latest/commands/sadd/) | ✅ | |
  | SREM | [https://redis.io/docs/latest/commands/srem](https://redis.io/docs/latest/commands/srem/) | ✅ | |
  | SCARD | [https://redis.io/docs/latest/commands/scard](https://redis.io/docs/latest/commands/scard/) | ✅ | |
  | SISMEMBER | [https://redis.io/docs/latest/commands/sismember](https://redis.io/docs/latest/commands/sismember/) | ✅ | |

# Performance

# Getting Started

# Architecture

## Core System Components

<p align="center">
  <img src="./assests/core_system_components.png"/>
  <br>
  <sub><em>Core components of Gopher Cache</em></sub>
</p>

#### TCP Server

The TCP server listens for incoming client connections on a configurable address and port. It acts as the entry point into the system and is responsible for receiving raw RESP-encoded commands and dispatching them for execution.

Internally, the server coordinates several core components: a command parser and a command queue. Each incoming command is parsed into a structured representation containing the command name and its arguments, then packaged into a command object and enqueued for execution.

The connection handler waits for the corresponding response and returns it to the client, preserving a synchronous request–response flow while allowing command execution to be decoupled from network I/O.

#### Command Queue

The Command Queue provides a thread-safe mechanism for decoupling command ingestion from command execution.

It acts as an intermediate buffer between the TCP server and the command executor, ensuring that commands are processed in a predictable order while allowing multiple client connections to submit commands concurrently.

#### Command Executor

The Command Executor is responsible for applying parsed commands to the underlying storage engine and producing protocol-compliant responses.

It synchronously processes commands from the Command Queue, dispatches them to the appropriate command handlers, and returns the result to the waiting client connection.

#### KV Store

The Storage Engine maintains an in-memory mapping of keys to typed values such as strings, lists, sets, and hashes. Each command handler function validates the expected type and either applies the requested mutation or returns a protocol-compatible error.

## Execution Flow

The following sequence diagram illustrates how a single client command is received, parsed, queued, executed, and responded to.

<p align="center">
  <img src="./assests/execution_flow_sequence_diagram.png"/>
  <br>
  <sub><em>Execution Flow Sequence Diagram</em></sub>
</p>

# Examples

#### Using Go

#### Using Python

#### Using CLI
