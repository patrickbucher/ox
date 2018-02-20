# ox

Go CLI for the Oxford Dictionary API

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
