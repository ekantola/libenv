# Libenv
[![Build Status](https://travis-ci.org/helkasko/libenv.svg?branch=master)](https://travis-ci.org/helkasko/libenv)

A Go library offering functionality for handling environmental variables.

This implementation is environment-agnostic and deals with the variables passed to it upon initialization or set later.

### Usage

A new instance can be invoked either by
- using the operating system's environment by calling `New` or
- using environment passed by the user by calling `NewFromMap`.

All alterations to the environment persist only in the execution context, meaning that the environment used in the instantiation (either passed or parsed) is not affected by the modifications made in runtime.

### Install

`go get github.com/helkasko/libenv`
