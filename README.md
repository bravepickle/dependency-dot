# dependency-dot
Generate entity relations graph using dot language notation

Using [GraphViz](http://www.graphviz.org/) you can visualize these entities as a graph in format you like from CSV files.

Application is written in [Golang](https://golang.org/). You can customize the way you want

## Installation

1. Clone repo
```bash
git clone https://github.com/bravepickle/dependency-dot
```
2. Change directory
```bash
cd dependency-dot/src
```
3. Build and install
```bash
go install -o depdot
```
4. Check application
```bash
$GOBIN/depdot -help
```
5. Done!

-- OR --

As an alternative you can just download executable binary file and just run it directly:
1. Download file
```bash
sudo wget https://github.com/bravepickle/dependency-dot/bin/depdot -O /usr/local/bin/depdot
```
2. Check application
```bash
depdot -help
```
3. Done!


## Usage

## TODO
- support grouping entities according to column
- customize output styles for entities

