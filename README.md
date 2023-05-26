# Text Modification Tool (In Development)

This tool will allow you to modify text based on specified rules. It will accept an input file containing the text to be modified, apply the specified modifications, and write the modified text to an output file.

## Plan

1. **Read the input file:** The tool will accept the input and output file names as command-line arguments. It will open and read the contents of the input file into a string variable. **Done and Tested - OK!****

2. **Parse the modifications:** A function will be defined to take the text as input and return the modified text. It will iterate over the text and search for the specified modifications. Regular expressions or string manipulation will be used to identify and extract the modification details. The appropriate modification will be applied based on the provided rules. **Done and Tested - OK!****

3. **Apply modifications:** Functions will be implemented for each modification type:
   - Convert hexadecimal to decimal: `ConvertHexToDecimal(hex string) string` **Done and Tested - OK!**
   - Convert binary to decimal: `ConvertBinaryToDecimal(binary string) string` **Done and Tested - OK!**
   - Convert to uppercase: `ConvertToUppercase(word string) string` **Done and Tested - OK!**
   - Convert to lowercase: `ConvertToLowercase(word string) string` **Done and Tested - OK!**
   - Capitalize a word: `CapitalizeWord(word string) string` **Done and Tested - OK!**
   - Apply case modification to a specific number of words: `ModifyWordsCase(text string, caseType string, wordCount int) string` **Done and Tested - `whitespacing` issue found with `,`**

   The corresponding modification function will be applied based on the parsed modifications.

4. **Write the modified text to the output file:** The tool will open the output file for writing and write the modified text to the output file.

5. **Test and validate:** Test cases will be created to cover different scenarios, including various modifications and edge cases. Unit tests will be written for the implemented functions. The correctness of the output will be validated against the expected results.

## Usage

To use this tool:

1. Clone the repository:

   ```shell
   git clone "Soon"
