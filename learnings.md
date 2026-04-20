# Learnings

I mainly created this with Go to learn.
Creating this with Laravel, with which I have the most programming experience, and thus with the Carbon date manipulation libraray would not have been a great challenge.

Besides, with the requirement of having to generate a CSV, I thought a CLI tool would be more than sufficient.

I have never created a full program with Go before, so wanted to grab this opportunity to do so.

What I have learned so far:

* Writing functions, if statements, switch case statements
* Different variable declaration styles `:=` and `var name type`.
* How to install new dependencies
* Working with several idiomatic libraries like `time`, `fmt` and `encoding/csv`
* That `if err := fooBar(); err != nil` is a preferred way of doing if with errors because it scopes the error to the if statement, and thus prevents leaking it outside.
