CRUD layer for the application.

Currently using `ent` for the storage layer. It's an ORM built by facebook.

Important that we don't let it leak out of this package though, don't want a headache if we decide we hate it.

Ent works by modifying the schemas in `ent/schema` and then running `go generate ./ent` to generate boilerplate.