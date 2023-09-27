using System;

class Example
{
    static void Main()
    {
        // Example of printing to the console
        Console.WriteLine("Hello, World!");

        // Basic math operations
        int num1 = 10;
        int num2 = 5;

        // Addition
        int sum = Add(num1, num2);
        Console.WriteLine($"Sum: {sum}");

        // Subtraction
        int difference = Subtract(num1, num2);
        Console.WriteLine($"Difference: {difference}");

        // Multiplication
        int product = Multiply(num1, num2);
        Console.WriteLine($"Product: {product}");

        // Division with error handling
        ResultAndError resultAndError = Divide(num1, num2);
        if (resultAndError.Error != null)
        {
            Console.WriteLine($"Error: {resultAndError.Error}");
        }
        else
        {
            Console.WriteLine($"Quotient: {resultAndError.Result}");
        }

        // Example of using a loop
        for (int i = 0; i < 5; i++)
        {
            Console.WriteLine($"Loop iteration {i}");
        }
    }

    // Function to add two integers
    static int Add(int a, int b)
    {
        return a + b;
    }

    // Function to subtract two integers
    static int Subtract(int a, int b)
    {
        return a - b;
    }

    // Function to multiply two integers
    static int Multiply(int a, int b)
    {
        return a * b;
    }

    // Function to divide two integers, returns an error for division by zero
    static ResultAndError Divide(int a, int b)
    {
        ResultAndError resultAndError = new ResultAndError();
        if (b == 0)
        {
            resultAndError.Error = "Division by zero";
        }
        else
        {
            resultAndError.Result = (double)a / b;
        }
        return resultAndError;
    }

    // Helper class to store both a result and an error
    class ResultAndError
    {
        public double Result { get; set; }
        public string Error { get; set; }
    }
}
