public class Example {
    public static void main(String[] args) {
        // Example of printing to the console
        System.out.println("Hello, World!");

        // Basic math operations
        int num1 = 10;
        int num2 = 5;

        // Addition
        int sum = add(num1, num2);
        System.out.println("Sum: " + sum);

        // Subtraction
        int difference = subtract(num1, num2);
        System.out.println("Difference: " + difference);

        // Multiplication
        int product = multiply(num1, num2);
        System.out.println("Product: " + product);

        // Division with error handling
        ResultAndError resultAndError = divide(num1, num2);
        if (resultAndError.error != null) {
            System.out.println("Error: " + resultAndError.error);
        } else {
            System.out.println("Quotient: " + resultAndError.result);
        }

        // Example of using a loop
        for (int i = 0; i < 5; i++) {
            System.out.println("Loop iteration " + i);
        }
    }

    // Function to add two integers
    public static int add(int a, int b) {
        return a + b;
    }

    // Function to subtract two integers
    public static int subtract(int a, int b) {
        return a - b;
    }

    // Function to multiply two integers
    public static int multiply(int a, int b) {
        return a * b;
    }

    // Function to divide two integers, returns an error for division by zero
    public static ResultAndError divide(int a, int b) {
        ResultAndError resultAndError = new ResultAndError();
        if (b == 0) {
            resultAndError.error = "Division by zero";
        } else {
            resultAndError.result = (double) a / b;
        }
        return resultAndError;
    }

    // Helper class to store both a result and an error
    public static class ResultAndError {
        public Double result;
        public String error;
    }
}
