# Hamming Distance Generator

## Author
Juan José Serrano

## Description
This program generates all possible binary strings of a given length that have a specific Hamming distance from the zero string (a string of all zeros).

The Hamming distance is the number of positions at which two strings differ. In this case, we generate all binary strings that have exactly 'h' ones (which represents the Hamming distance from a string of all zeros).

## Features
- Concurrent implementation using goroutines
- Input validation for string length and Hamming distance
- Efficient recursive algorithm
- Thread-safe result collection using mutex

## Usage
1. Run the program
2. Enter the desired length of binary strings when prompted
3. Enter the desired Hamming distance when prompted
4. The program will display all valid binary strings that meet the criteria

## Example

Enter the length of the binary strings: 4
Enter the desired Hamming distance: 2

Found 6 binary strings:
Chain 1: 1100
Chain 2: 1010
Chain 3: 1001
Chain 4: 0110
Chain 5: 0101
Chain 6: 0011

## Implementation Details
The program uses a recursive backtracking algorithm with concurrent execution using goroutines.

The main algorithm can be found in the `hammingDistanceRec` function.


## Limitations
- The Hamming distance must not exceed the length of the strings
- Due to the concurrent nature of the implementation, the order of results may vary between executions


