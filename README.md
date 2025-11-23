# kami

## Description
A small but capable router library for the go programming language that speaks JSON.

## Objectives
The library is primarily aimed at microservices that back a frontend applications consuming JSON via JavaScript's Fetch API.

## Philosophy
The library should be small enough that the code and the tests can be consumed as the documentation.


## Usage

### Paths

- parameters are defined with a leading colon ":"
    - the router disallows paths prefixes followed by a different parameter name for example the registering both these paths would lead to a panic.

    "/foo/bar/:buzz"
    "/foo/bar/:bazz"

- wildcards are defined with a leading asterisk "*"