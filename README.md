# Introduction
`pol` is a CLI tool to automate the process of adding images to a problem on [Polygon](https://polygon.codeforces.com/). It takes in a directory name as a parameter, traverses that directory to scan all the images, then creates a `TeX` block of all the images in a sorted order. It also copies the output to your clipboard, enabling you to directly paste the contents in your desired location. It also has autocomplete feature to ensure speed/accuracy while typing directoy names.

A sample output may look like : 


```
\begin{center}
\includegraphics{image-1.png}
\includegraphics{image-2.png}
\includegraphics{image-3.png}
\end{center}
```

---

# Installation Manual
1. Download and install **Go** : [Link](https://golang.org/doc/install)
2. Add the `go` binary to your `PATH`.
3. Set your `GOPATH`
4. Add your workspace's `GOPATH/bin` directory to your `PATH`.

You can perform the last 3 steps by adding these 3 lines at the end of your `~/.profile`.

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Refresh the shell

```bash
source ~/.profile
```

Install the binary

```bash
go install github.com/variety-jones/pol@latest
```

Enable auto completion of directories

```bash
echo "complete -A directory pol" > pol-completion
sudo mv pol-completion /etc/bash_completion.d/pol-completion
```

Install `xclip` to copy the contents to clipboard

```bash
sudo apt-get update
sudo apt-get install xclip
```

---

# Usage
From any directory, type

```bash
pol /relative/path/to/other/directory
```

You can keep pressing tabs for autocompletion. The resulting output should automatically be copied to the clipboard.

---

# Examples
Create the statement file from the images in the current directory.
```bash
pol .
```

Create the statement file from the images of a different directory
```bash
pol ../directory-1/directory-2
```
