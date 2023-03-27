// Package customsort provides types and methods to handle the `sort` query
// param. It provides two types: `Sort` and `SortMap`. The `Sort` type is a
// string, and the `SortMap` type is a map of `fieldName:order` pairs.
//
// SortMap provides a custom marshaler which helps Echo to properly convert the
// `sort` query to a `SortMap` type. In this case, it not only converts but also
// validates that. The downside of using it is that it can only be used with
// Echo's `Bind` method.
//
// The `Sort` type is more powerful and flexible, it's
// framework agnostic, can be easily extended adding more formats. It's used
// under-the-hood by the `SortMap` type.
package customsort
