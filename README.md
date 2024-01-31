# Simple Go Bank Account Project

This is a basic Go project showcasing fundamental concepts such as interfaces, structs, and error handling. The project simulates a simple banking system with different types of accounts.

## Overview

The project consists of the following components:

- `main.go`: The main program file where bank accounts are created, operations are performed, and results are displayed.
- `account.go`: Defines various account types like `Account`, `LoanerAccount`, `PersonalAccount`, and `BusinessAccount` along with associated operations.

## Key Concepts

### Structs
- Utilizes structs to represent different account types, each with its own set of properties.

### Interfaces
- Demonstrates the use of interfaces with the `User` interface, which is implemented by various account types.

### Error Handling
- Implements a custom error (`ErrCreateBusinessAccount`) for specific error scenarios, such as attempting to add a non-business account to a legal person's (`PJ`) account list.

## How to Run

To run the project, use the following command in your terminal:

```bash
go run main.go


