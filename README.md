## grav

This is a simple, silly CLI for the gravatar API. You can install it like this:

```
go get -u github.com/petderek/go-gravatar/grav
go install github.com/petderek/go-gravatar/grav
```

For now, it lets you answer two common questions:

1. Does `derek@example.com` have a gravatar?

```
grav -check derek@example.com && echo YES || echo NO
```

2. What is the gravatar for `derek@example.com`?

```
grav -save derek.jpeg derek@example.com
```

![derek@example.com](https://gravatar.com/avatar/631f0aac0a664c033945d4cb2573f931)

3. What if I'm using a different domain?

```
grav -check -domain http://cdn.libravatar.org -sha256 derek@example.com
```
