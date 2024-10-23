# GoCache

Welcome to **GoCache**—an elegant in-memory caching solution crafted in Go (Golang).

## Overview

At its core, GoCache implements a Least Recently Used (LRU) caching mechanism. This approach ensures that the most frequently accessed data remains readily available, while less frequently accessed data is gracefully evicted when necessary. 

### Features

- **LRU Mechanism**: Automatically manages memory by removing the least recently accessed entries.
- **Thread-Safe**: Built with concurrent access in mind, leveraging goroutines to ensure data integrity.
- **Simplicity**: A clean and intuitive API for setting, retrieving, and deleting cached values.

## Getting Started

1. **Clone the Repository**:
   git clone https://github.com/sameerihs/GoCache.git
   cd GoCache
