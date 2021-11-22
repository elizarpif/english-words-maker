# english-words-maker

### 1. Create loto game files and vocabulary file
```go
words, err := jsonwords.GetWords()
if err != nil {
    log.Fatal(err)
}

err = maker.Vocabulary(words)
if err != nil {
    log.Fatal(err)
}

err = maker.CreateGames(words, 3)
if err != nil {
    log.Fatal(err)
}
```

### 2. Get random word for activity game
```go
words, err := jsonwords.GetWords()
if err != nil {
    log.Fatal(err)
}

mapW := maker.GetMap(words)
sc := bufio.NewScanner(os.Stdin)

for key, value := range mapW {
    fmt.Printf("generated %s - %s\n", key, value)
    sc.Scan()
}
```

Example
```shell
>> Activity started

>> generated "battle" - [бой]

>> generated "dimension" - [измерение]
```


### 2. StoryTeller
```go
	err := maker.StoryTeller(3)
	if err != nil {
		log.Fatal(err)
	}
```

Example
```shell
>> StoryTeller for 3 people started

>> generated: (vanity, intrinsic, obtain, cause)

>> generated: (subsidy, to denote, perspective, inflation)

>> generated: (hypocritically, fringe, manipulation, minimum)
```