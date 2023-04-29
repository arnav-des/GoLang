package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
)

func PrettyString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}

func main() {
    fruitJSON := `{"name": "Strawberry", "color": "red"}`
    res, err := PrettyString(fruitJSON)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(res)
}
