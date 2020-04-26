# starcrossedGen

> I made this project for my personal use, therefore, some features may be missing. I will be adding more features as my need.

`sg` (*starcrossedGen*) is a meta-build system that I use for my personal projects, it generates build files for Ninja.

## Getting started

### Dependencies

+ [The Go Programming Language](https://golang.org/)

### Building

First of all, clone the repository.

```bash
git clone https://github.com/starcrossedRadio/sg.git
```

Then compile.

```bash
go build
```

The executable will be in the folder where you compiled it.

## Running the example

```bash
cd example
mkdir out
../sg
ninja -C out
```
```bash
./out/HelloWorld
```

## Contributing

Contributions through Pull Requests are welcome.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
