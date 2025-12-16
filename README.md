# Rush01-Go

A Go implementation of the classic "Skyscraper" (Rush01) puzzle solver, inspired by exercises from 42.

## Overview

This project demonstrates solving the Rush01 (Skyscraper) puzzle using Go. The puzzle consists of filling an NxN grid (4x4, 5x5, or 6x6) with numbers 1-N so that:

- Each row and column contains each number exactly once.
- Clues around the grid indicate how many skyscrapers are visible from that direction.

## Features

- Command-line interface for inputting clues.
- Supports 4x4, 5x5, and 6x6 grids.
- Step-by-step board printing for debugging and learning.
- Modular code structure for clarity and maintainability.

## Usage

### Build and Run

From the `ex00` directory, run:

```sh
go run . <clues>
```

Or, to specify files explicitly:

```sh
go run rush01.go solver.go <clues>
```

### Input Format

Clues must be provided as a **single string of digits** in this order:

1. Top clues (left to right)  
2. Bottom clues (left to right)  
3. Left clues (top to bottom)  
4. Right clues (top to bottom)  

**Accepted lengths:**
- 16 digits for 4x4
- 20 digits for 5x5
- 24 digits for 6x6

**Example for 4x4:**

```sh
go run . 4132234214321234
```

**Example for 5x5:**

```sh
go run . 33212133024321332132
```

**Example for 6x6:**

```sh
go run . 234321234321234321234321
```

### Example Output

```
     4   1   3   2 
   +---+---+---+---+
 2 | 1 | 2 | 3 | 4 | 3
   +---+---+---+---+
 2 | 4 | 3 | 2 | 1 | 2
   +---+---+---+---+
 3 | 2 | 1 | 4 | 3 | 4
   +---+---+---+---+
 1 | 3 | 4 | 1 | 2 | 2
   +---+---+---+---+
     4   3   2   1 
```

## Learning Goals

- Practice Go syntax and idioms.
- Implement backtracking algorithms.
- Work with slices, functions, and modular code.

## License

MIT License
