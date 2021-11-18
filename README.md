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
>> generated esthetically - эстетично
>> generated cause - причина
...
>> generated striking - поразительный, заметный
```