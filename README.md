# ox

Go CLI for the Oxford Dictionary API

**NOTE: The Oxford Dictionary does support far less languages then I thought. So contrary to the initial plan, the tool will primarly serve as an English Dictionary**

## Setup

1. Create your [Oxford Dictionaries Developer Account](https://developer.oxforddictionaries.com/)
2. Find your [API credentials](https://developer.oxforddictionaries.com/admin/applications)
3. Modify `api.go.template`, insert your credentials and save it as `api.go`
4. Build the CLI tool: `go build cmd/ox.go`

## Usage (*currently implemented*)

Look up the word "beer" in the English dictionary:

`ox -lang en beer`

English is the default language, so you can just type:

`ox beer`

Look up the word "trabajo" in the Spanish dictionary:

`ox -lang es trabajo`

## Usage (*to be implemented*)

### Dictionary Lookup

Look up the word "television" in the English dictionary:

`ox dict -lang en television`

Look up the word "достопримечательность" in the Russian dictionary:

`ox dict -lang ru достопримечательность`

Short syantax:

```
ox dict -l en television
ox dict -l ru достопримечательность
```

### Translation

Translate the English word "beer" to German:

`ox trans -from en -to de beer`

Translate the French word "chien" to Russian:

`ox trans -from fr -to ru chien`

Short syntax:

```
ox trans -f en -t de beer
ox trans -f fr -t ru chien
```

### Misc

List available languages:

`ox langs`
