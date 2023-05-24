# Text Modification Tool

The Text Modification Tool is a command-line program written in Go that allows you to modify text based on certain rules and patterns. It provides various functionalities to transform and manipulate the input text according to specific patterns and criteria.

## Features

The Text Modification Tool supports the following features:

- Conversion of hexadecimal and binary numbers:
  - Replaces instances of (hex) with the decimal version of the preceding hexadecimal number.
  - Replaces instances of (bin) with the decimal version of the preceding binary number.

- Text case modification:
  - Converts the word before (up) to uppercase.
  - Converts the word before (low) to lowercase.
  - Converts the word before (cap) to capitalized form.

- Customized case modification:
  - If a number appears next to (up), (low), or (cap), the program modifies the specified number of preceding words accordingly.

- Punctuation handling:
  - Properly formats punctuation marks (. , ! ? : ;) by placing them close to the previous word and with a space separating them from the next word.
  - Handles groups of punctuation marks (...) or (!?).
  - Preserves the placement of single quotation marks (' ') around words or phrases.

- Indefinite article correction:
  - Changes the indefinite article 'a' to 'an' if the next word begins with a vowel (a, e, i, o, u) or an 'h'.

## Usage

To use the Text Modification Tool, follow these steps:

1. Install Go on your system if it is not already installed.
2. Clone the repository or download the source code.
3. Open a terminal and navigate to the project directory.
4. Run the following command to build the program:
go build

lua
Copy code
5. Run the program with the desired input and output files:
./text-modification-tool <input_file> <output_file>

vbnet
Copy code
Replace `<input_file>` with the path to the file containing the text to be modified and `<output_file>` with the desired path to save the modified text.
6. The program will process the input file, apply the specified modifications, and save the result in the output file.

## Language and Dependencies

The Text Modification Tool is written in the Go programming language (Golang). It leverages Go's standard library for file I/O, regular expressions, string manipulation, and other operations.

No external dependencies are required to run the program.
