## Go Diceware Generator

Diceware implementation in Go.  
Wordlist generated from the [EFF's](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) "long" list.

### To Build From Source

```bash
git clone https://github.com/abdullahrehmat/go-diceware  
cd ./Go-Diceware  
go build -ldflags "-ws" main.go  
chmod +x ./main  
./main  
```