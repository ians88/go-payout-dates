# Learnings

I mainly created this with Go to learn. I have never created a full program with Go before, so wanted to grab this opportunity to do so.

What I have learned so far:

* Writing functions, if statements, switch case statements
* Different variable declaration styles `:=` and `var name type`.
* How to install new dependencies
* Working with several idiomatic libraries like `time`, `fmt` and `encoding/csv`
* That `if err := fooBar(); err != nil` is a preferred way of doing if with errors because it scopes the error to the if statement, and thus prevents leaking it outside.
* That compiling into a binary is really easy with Golang
