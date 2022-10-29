# Study Golang Basic
This will be the first walk thru while studying Golang 😁

Most of this study will be made using VSCode, in a remote Ubuntu machine.

## Golang Instalation

If you don't have golang installed already (`golang --version` returns nothing)
Then you'll have to install golang

```bash
# Remove any previous Go installation
mkdir -p ~/tmp
cd ~/tmp
# you can get this wget link from: https://go.dev/doc/install
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
rm -f ~/tmp/go1.19.2.linux-amd64.tar.gz
# hopefully you're using zsh too, otherwise change the file to .bashrc =]
grep -qxF 'export PATH=$PATH:/usr/local/go/bin' ~/.zshrc || echo '\n\n# Golang Path\nexport PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
go version
```
> [source](https://go.dev/doc/install) - plus some automations

## Hello World

Go to your project folder, in my case: `~/code/study-golang-basic`
```bash
cd ~/code/study-golang-basic
go mod init example/hello
```

Create a file `hello.go` and add the code:  
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Now let's see if that really works:
```
go run .
```
If in vscode, install the golang extension, plus the gopls:  
![Alert in VSCode](img/vscode-gopls.png)  
I've installed in the end:
- gopls
- goimports
- go-outlie

### Run direct from VSCode:
With the `hello.go` file open, you can run the debugger. ⌨️`F5`  
You can also create a launch.json, for running with F5 from any file  
<details>
  <summary>My `launch.json` here</summary>

  ```js
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            // "program": "${fileDirname}"
            "program": "hello.go"
        }
    ]
}
  ```
</details>

## Let's add some extensions/modules!
> [source](https://go.dev/doc/tutorial/getting-started#call)

Search the next nice package in [pkg.go.dev](pkg.go.dev)
In our case, "quotes"

Changes to `hello.go`
```go
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

```
Add new module requirements and sums.
```bash
go mod tidy
```

## Creating a new Module

Part of this study continues in [MadMod ](https://github.com/madpin/madmod)
I've created the module there, and will use it here.

We'll change the `hello.go` file, new content:
```go
package main

import (
	"fmt"

	"github.com/madpin/madmod"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Opt())
	fmt.Println(madmod.Hello("test"))
}
```

We'll also need to change the target of our module, to get from local:
```bash
go mod edit -replace github.com/madpin/madmod=../madmod
```

Now, the last sync before the run:
```bash
go mod tidy
```
> Once ran, you can edit the module code, that it's auto update for the hello.