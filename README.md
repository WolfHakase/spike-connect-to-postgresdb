# spike-connect-to-postgresdb
This is a test project to see how to interact with a running postgresdb from Go code.

## What is a spike?
According to Wikipedia: 

`A spike is a product development method originating from extreme programming that uses the simplest possible program to explore potential solutions. It is used to determine how much work will be required to solve or work around a software issue. Typically, a "spike test" involves gathering additional information or testing for easily reproduced edge cases. The term is used in agile software development approaches like Scrum or Extreme Programming.`

In these spikes no use is made of TDD. In fact, the code is largely untested. 

The code isn't necessarily clean. 

It is merely a tool to understand the subject, so that the actual production code can more easily be tested and made clean.

An added benefit to spikes is that they don't contain anything worth stealing, therefore they're perfect to build a portfolio with.

## Integration tests
Integration tests test the integration between components. Nothing is mocked, as opposed to unit tests. Both offer their own merits.

This project has also been used to test how to write integration tests for the database, as well as environment variables.

## How to run
To run this project you don't have to do anything special.

From Goland go to the `main.go` and use the hotkey `ctrl + shift + f10`.

Or if you don't use a fancy IDE just run it like you would run any other Go script.