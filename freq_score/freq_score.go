package freq_score
//package main

//import ("fmt"; "strings")
import ("strings")

//func main(){
func init(){
    /*text := []byte("a1b'")
    score := scoring_system(text)
    fmt.Printf("%d\n", score)*/
}

func Score(txt []byte)int {
    tmp_score := 0
    tmp_score = letters(txt, tmp_score)
    tmp_score = symbols(txt, tmp_score)
    return tmp_score
}

func letters(letter_bytes []byte, letter_score int)int{
    key_letters := map[string]int{
        "E": 5, "T": 5, "A": 5, "O": 5, "N": 5,
        "e": 5, "t": 5, "a": 5, "o": 5, "n": 5,
        "R": 4, "I": 4, "S": 4, "H": 4, "D": 4,
        "r": 4, "i": 4, "s": 4, "h": 4, "d": 4,
        "L": 3, "F": 3, "C": 3, "M": 3, "U": 3,
        "l": 3, "f": 3, "c": 3, "m": 3, "u": 3,
        "G": 2, "Y": 2, "P": 2, "W": 2, "B": 2,
        "g": 2, "y": 2, "p": 2, "w": 2, "b": 2,
        "V": 1, "K": 1, "J": 1, "X": 1, "Q": 1,
        "v": 1, "k": 1, "j": 1, "x": 1, "q": 1,
        "Z": 1, "z": 1,
    }
    letter_score = calculate(key_letters, letter_bytes, letter_score)
    return letter_score
}

func symbols(symb_bytes []byte, symbol_score int)int{
    key_symbols := map[string]int{
        "'":    -10, "\"":   -10, "!":    -10, "@":    -10,
        "#":    -10, "$":    -10, "%":    -10, "¨":    -10,
        "&":    -10, "*":    -10, "(":    -10, ")":    -10,
        "-":    -10, "=":    -10, ",":    -10, ":":    -10,
        ";":    -10, "?":    -10, "{":    -10, "}":    -10,
        "[":    -10, "]":    -10, "<":    -10, ">":    -10,
        "/":    -10, "^":    -10, "~":    -10, "|":    -10,
        ".":    -10, "\\":   -10, "`":    -10, "´":    -10,
        " ":      5,
        /*"":   -10, "": -10, "": -10, "": -10,
        "": -10, "": -10, "": -10, "": -10,
        "": -10, "": -10, "": -10, "": -10,
        "": -10, "": -10, "": -10, */
    }
    symbol_score = calculate(key_symbols, symb_bytes, symbol_score)
    return symbol_score
}

func digraphs(letter_bytes []byte, digraph_score int)int{
    key_digraph := map[string]int{
        "TH": 20, "ER": 20, "ON": 20, "AN": 20, "RE": 20,
        "th": 20, "er": 20, "on": 20, "an": 20, "re": 20,
        "HE": 20, "IN": 20, "ED": 20, "ND": 20, "HA": 20,
        "he": 20, "in": 20, "ed": 20, "nd": 20, "ha": 20,
        "AT": 20, "EN": 20, "ES": 20, "OF": 20, "OR": 20,
        "at": 20, "en": 20, "es": 20, "of": 20, "or": 20,
        "NT": 20, "EA": 20, "TI": 20, "TO": 20, "IT": 20,
        "nt": 20, "ea": 20, "ti": 20, "to": 20, "it": 20,
        "ST": 20, "IO": 20, "LE": 20, "IS": 20, "OU": 20,
        "st": 20, "io": 20, "le": 20, "is": 20, "ou": 20,
        "AR": 20, "AS": 20, "DE": 20, "RT": 20, "VE": 20,
        "ar": 20, "as": 20, "de": 20, "rt": 20, "ve": 20,
    }
    digraph_score = calculate(key_digraph, letter_bytes, digraph_score)
    return digraph_score
}



func calculate(_map map[string]int, long_bytes []byte, mid_score int)int{
    for char,value := range _map{
        //fmt.Printf("%s - %d\n", x, y)
        count := strings.Count(string(long_bytes), char)
        mid_score += count*value
    }
    return mid_score
}
