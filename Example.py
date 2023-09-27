def main():
    # Example of printing to the console
    print("Hello, World!")

    # Basic math operations
    num1 = 10
    num2 = 5

    # Addition
    sum_result = add(num1, num2)
    print(f"Sum: {sum_result}")

    # Subtraction
    difference = subtract(num1, num2)
    print(f"Difference: {difference}")

    # Multiplication
    product = multiply(num1, num2)
    print(f"Product: {product}")

    # Division with error handling
    quotient, error = divide(num1, num2)
    if error is not None:
        print(f"Error: {error}")
    else:
        print(f"Quotient: {quotient}")

    # Example of using a loop
    for i in range(5):
        print(f"Loop iteration {i}")

# Function to add two numbers
def add(a, b):
    return a + b

# Function to subtract two numbers
def subtract(a, b):
    return a - b

# Function to multiply two numbers
def multiply(a, b):
    return a * b

# Function to divide two numbers, returns an error for division by zero
def divide(a, b):
    if b == 0:
        return None, "Division by zero"
    return a / b, None

if __name__ == "__main__":
    main()