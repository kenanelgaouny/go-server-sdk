// Package testhelpers contains types and functions that may be useful in testing SDK functionality.
//
// Currently, this consists of only the storetest subpackage, which is meant to be used by any
// implementation of a persistent data store. If you are writing your own database integration, use this
// test suite to ensure that it is being fully tested in the same way that all of the built-in ones are
// tested.
//
// The APIs in this package and its subpackages are supported as part of the SDK. Purely internal test
// helpers that are likely to change when SDK implementation details change should not be in this package,
// but instead in internal/sharedtest.
package testhelpers
