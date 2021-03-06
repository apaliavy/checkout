# Checkout Kata

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)  
- [Installation](#installation)
- [Testing](#testing)  
- [Application](#application)

## Introduction
Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. In addition, some items are multi- priced: buy n of them, and they’ll cost you y pence. For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they’ll cost you 130. In fact the prices are:

| SKU           | Unit Price    | Special Price |
| ------------- | ------------- | ------------- |
|       A       |       50      |   3  for 130  |
|       B       |       30      |   2  for 45   |
|       C       |       20      |               |
|       D       |       15      |               |

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95). The pricing changes frequently, so pricing should be independent of the checkout.

The interface to the checkout could look like:

```java
interface ICheckout
{
  void Scan(string item);
  int GetTotalPrice();
}
```

## Prerequisites

Ensure you have Golang installed. [Here](https://golang.org/doc/install) you can find official manual which describes how to do it step by step.

## Installation 

Download repository: 

```
git clone git@github.com:apaliavy/checkout.git
```

Use Makefile to install all dependencies: 

```
cd checkout 
make deps
```

Now you can run the code:
```
make run
```

The command above builds a binary in the /bin folder for you and runs it. There are some useful env variables available:

- OS (`darwin` by default) - allows you to specify your host OS; 
- SERVICE_NAME (`checkout` by default) - allows you to specify your binary name;
- LOG_LEVEL (`debug` by default) - allows you to specify default logging level;

Full example: 

```
OS=linux SERVICE_NAME=my-checkout LOG_LEVEL=INFO make run
```

Makes a binary(`my-checkout`) for `linux` under your bin/ folder and runs it with `INFO` logging level. 

To run a static checks for you code use this command:
```
make lint 
```

## Testing 

To run the whole set of unit tests run:

```
make test
```

## Application

### `/cmd`

Main applications for this project.

The directory name for each application match the name of the executable we want to have (e.g., `/cmd/checkout`).

### `/app`

Private application and library code. This is the code we don't want others importing in their applications or libraries. 

Here we have the following packages: 

#### `/app/discount`

This package contains information about special offers, special offers collection and methods to works with them.

#### `/app/price`

This package contains price calculation engine. `Calculator` calculates the total price, including information about special offers and regular prices.

#### `/app/stock` 

This package contains such entities like `Product` (the regular product, with non-special price) and collection to work with a set of product. Also, here we have a definition of `SKU` type.

### `/bin`

Place for project binaries (generated using `make build`)

### `/testing`

Place for generated mock files. Use `make mocks` to refresh content in this folder.

### `/tools`

A special folder, needed to import `counterfeiter` library (used to generate mock files).