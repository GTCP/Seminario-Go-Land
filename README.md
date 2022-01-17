# Seminario-Go-Land

This folder contains go exercises, it also contains a final exercise of the seminar. This has a function called Convertor, which is passed a string and returns a structure and an error if there is none. To do this, we create a new structure of the Result type, then we assign the attributes of said structure defined parts of the string, checking that these parts meet all the conditions:

That the variable type is 2 characters letter, either lowercase or uppercase.
That the Length is 2 int numbers and that this number matches the number of characters in the value
Let the Value be numeric or alphabetic characters but not both. Failure to comply with these conditions will notify the error on the screen.
In our main_test we run a json of several cases in which the converter can fail and through a for we print all the cases. If there is an error, it will be notified.

The test reaches 89.1% coverage.
