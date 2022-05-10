![](https://tuchuang-1300339532.cos.ap-chengdu.myqcloud.com/img/hackbox.png)


<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://goreportcard.com/badge/github.com/WAY29/hackbox"><img src="https://goreportcard.com/badge/github.com/WAY29/hackbox"></a>
<a href="https://github.com/WAY29/hackbox/releases"><img src="https://img.shields.io/github/release/WAY29/hackbox"></a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#running">Running</a> •
<a href="#notes">Notes</a>
</p>

hackbox is an lightweight and easy to use toolbox for hackers, it is designed to organize and uniformly call your favorite command line tools.

# Features
![](https://tuchuang-1300339532.cos.ap-chengdu.myqcloud.com/img/20220510094809.png)

- Highly **interactive command line framework** powered by [c-bata/go-prompt](https://github.com/c-bata/go-prompt)
- Easy to use and intelligent **automatic completion**
- **Rich and self-explanatory default configuration files**, so you can easily customize it
- **Built-in several command line tools for hackers** (in comming...)
- Clear arguments settings, no need to read annoying help documents, so you can **call tools faster and more conveniently**
- Save the execution result of the tool which can be **used as input to another tool or easily exported**
- Command line parameters **support validate** powered by [go-playground/validator](https://github.com/go-playground/validator)

# Installation
## go get
hackbox requires go1.17 to install. Run the following command to get hackbox:
```go
go install -v github.com/WAY29/hackbox/cmd/hackbox@latest
```

## static releases
Download hackbox from [Releases](https://github.com/WAY29/hackbox/releases)

# Usage
Just run hackbox, and you will get an interactive hackbox shell. Here are hackbox help documents:
```
Usage of hackbox:
  -nc
        Print without color
  -p string
        Custom tool path, default will load from ./tools.toml -> $HOME/.config/hackbox/tools.yaml
  -q    Run hackbox without banner
```

# Running
You can run some intuitive commands, such as `cd`, `ls`, `sh`, `exit`
- `cd <tool directory>` change tool directory
- `ls` List tools and subdirectories in this directory
- `sh <command>` Run local shell command
- `exit` Exit hackbox

When you find the tool you want, `use` it immediately, then you can `show` or `set` or `unset` arguments, then just `run` it.
- `use <tool name>` use specified tool
- `set <arg name> <arg value>` set tool argument
- `unset <arg name>` unset tool argument
- `show [arg name]` Show tool argument(s)
- `run [as <output name>] [bg]` Run tools [in background]

Then you can `output` the result or `export` it, before this, maybe you want to `filter` it or `merge` other result
- `output [output name]` show output(s)
- `filter <output name> <filter> [as <new output name>]` filter output by filters **link/email/date/time/phone/ip/md5/sha1/sha256**
- `merge <output1 name> <output2 name> as <output name>` merge two outputs as one
- `export <output name> [export path, default <output name>.txt]` export output as file

By the way, you can `setoutput` or `unsetoutput`
- `setoutput <output name> <output value / filepath>` set output from input or file
- `unsetoutput <output name>` unset output

# Notes
- It only took me three days to write this project, so there may be some mistakes, welcome issues.





