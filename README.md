# Unexpected Zero Values in Go Maps

This repository demonstrates a common, yet subtle, bug in Go involving the implicit zero values returned when accessing non-existent keys in maps.  This can lead to hard-to-find bugs if not handled correctly.

## The Bug

Go maps will return the zero value for a given type if you try to access a key that does not exist. This can lead to unexpected behavior, especially in situations where you expect an error to be thrown or a different default value. The example demonstrates how accessing a non-existent key returns 0 instead of producing an error or panic. This behavior is often unexpected by developers coming from languages where such an operation would throw an exception.

## The Solution

The solution involves explicitly checking if a key exists before accessing it using the `_, ok := map[key]` syntax. This provides more robust error handling and makes the code's intent clear.
