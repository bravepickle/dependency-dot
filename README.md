# dependency-dot
Generate entity relations graph using dot language notation

Using [GraphViz](http://www.graphviz.org/) you can visualize these entities as a graph in format you like from CSV files.

Application is written in [Golang](https://golang.org/). You can customize the way you want

## Installation

```bash
# 1. Clone repo
git clone https://github.com/bravepickle/dependency-dot

# 2. Change directory
cd dependency-dot/src

# 3. Build and install
go install -o depdot

# 4. Check application
$GOBIN/depdot -help

# 5. Done!
```

###-- OR --

As an alternative you can just download executable binary file and just run it directly:

```bash
# 1. Download file
sudo wget 'https://github.com/bravepickle/dependency-dot/blob/master/bin/depdot?raw=true' -O /usr/local/bin/depdot

# 2. Check application
depdot -help

# 3. Done!
```

## Usage
### Prerequisites
1. CSV file' should titles at first row and named accordingly to help notes of depdot file OR flags should be passed later on to executable script (depdot)
2. CSV file should be comma-separated and values should properly be escaped\wrapped, formatted
3. GraphViz should be installed

### Basic
```bash
depdot -o /tmp/out.dot test.csv
dot -Tpng /tmp/out.dot > /tmp/out.png
```
Will output dot file to /tmp/out.dot and will be generated PNG file to /tmp/out.png by next command using GraphViz

As an alternative one can execute
```bash
depdot test.csv | dot -Tpng >/tmp/out.png
```

Also you can directly update dot file to adjust it to your liking (ad groups, labels, styling etc.). See GraphViz [reference](http://www.graphviz.org/pdf/dotguide.pdf) and [attributes guide](http://www.graphviz.org/content/attrs) for more info


## TODO
- support grouping entities according to column
- customize output [styles](http://www.graphviz.org/content/attrs) for entities.
- styles for references

## Notes
- ensure that CSV files are generated correctly and special symbols (doublequotes, commas etc.) are properly escaped
- pay attention to warnings and errors during script execution

