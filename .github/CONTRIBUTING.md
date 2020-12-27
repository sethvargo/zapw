# Contributing

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.


## Please submit failing tests

If you've found a bug, please consider submitting a **Pull Request** that adds a
test to showcase the failure. This ensures the problem is fixed and remains
fixed in the future.

To submit a test, either create a new file in `pkg/testdata/src/defaults` or add
to an existing one.

If zapw is reporting an error when it shouldn't be, add the example code with
no comments. Go's analysis tester will fail if a failure is reported that was
not marked as expected.

```go
func testThing() {
  X.Y()
}
```

If zapw is **not** reporting an error when it should be, add the example code
with a comment where the error should be reported.

```go
func testThing() {
  x.Y() // want `zap.SugaredLogger should x,y,z...`
}
```

You can also [learn more about Go analysis testing](https://pkg.go.dev/golang.org/x/tools/go/analysis/analysistest).


## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.
