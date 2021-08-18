# Plugin Template
A plugin-template for MacCleanup.


## Usage
In order to create your own plugin you'll just need to use this template to create a new repository.

## Documentation
Use The `DeleteFiles` method within the `src/plugin.go` file like so:

```go
func (g plugin) Cleanup() {
	...
	DeleteFiles("~/DIR_YOU_WANT_TO_BE_CLEANED") // DIR_YOU_WANT_TO_BE_CLEANED is the path to the directory you want to be cleaned
    ...
    // Add more here
}
```

## Compile your plugin
To compile your plugin locally simply run:

```bash
go build -buildmode=plugin -o dist/plugin.so src/plugin.go
```

You now have your `plugin.so` file within the `dist` directory.
To use your plugin just copy your `plugin.so` file into `~/.mac-cleanup/plugins/YOUR_PLUGIN/`.
Run `mac-cleanup clean` to make use of any plugins in this directory.

## Contributing
Contributions are always welcome!
Please adhere to this project's `code of conduct`.

## License
This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/) license.

## Authors
- [@fwartner](https://www.github.com/fwartner)
