# Defer in loop

When you use a `defer` in a loop, you should wrap the actions in the loop as a function, otherwise:
- The `defer` functions will not be executed after every iteration ends, which is not expected by most coder.
- Resource leaks may occur.

As the test results show in [defer_in_loop_test.go](defer_in_loop_test.go)


## References
- [StackOverflow-Proper-way-to-release-resources-with-defer-in-a-loop](https://stackoverflow.com/questions/45617758/proper-way-to-release-resources-with-defer-in-a-loop)