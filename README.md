![Go Version](https://img.shields.io/badge/go-1.25+-blue)
![OS](https://img.shields.io/badge/OS-Windows,Linux,MacOS-orange)
![License](https://img.shields.io/badge/license-MIT-green)

# Table of Content

- [About](#about)
- [Features](#features)
- [Performance](#performance)
- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [Configuration](#configuration)
- [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)

# About

Gopher Cache is a small in-memory database. It was built with the intention to mock a Redis database, so that it behaves in a similar way as the real one from the user/client perspective.
The source code of the project is fully written in Go. The initial goal of this project was to build a mini copy of the Redis database in order to increase understanding of both
the Redis internals and the Go programming language. However, everybody is welcome to check the code, propose improvements, create their own forks/versions, and use the source code
for their own purposes.

# Features

- **RESP Command Parser:** recursive descent parser, able to parse RESP-style array of bulk strings
- **TCP Server:** a server for accepting and queuing the commands
- **FIFO Command Execution:** commands are executed in first-in-first-out order as they arrive to the server
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

**1. TCP Server:**
**2. Command Queue:**
**3. Command Executor:**
**4. KV Store:**

## Classes

## Execution Flow
