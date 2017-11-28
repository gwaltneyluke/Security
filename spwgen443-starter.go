////////////////////////////////////////////////////////////////////////////////
//
//  File           : spwgen443.go
//  Description    : This is the implementaiton file for the spwgen443 password
//                   generator program.  See assignment details.
//
//  Collaborators  : Luke Gwaltney, Jason Ling, Albert Wilson, Brian Young
//  Last Modified  : **TODO**: FILL ME IN
//

// Package statement
package main

// Imports
import ( 
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
	"github.com/pborman/getopt"
	// There will likely be several mode APIs you need
)

// Global data
var patternval string = `pattern (set of symbols defining password)

        A pattern consists of a string of characters "xxxxx",
        where the x pattern characters include:

          d - digit
          c - upper or lower case character
          l - lower case character
          u - upper case character
          w - random word from /usr/share/dict/words (or /usr/dict/words)
              note that w# will identify a word of length #, if possible
          s - special character in ~!@#$%^&*()-_=+{}[]:;/?<>,.|\

        Note: the pattern overrides other flags, e.g., -w`

const LOWERCASE string = "abcdefghijklmnopqrstuvwxyz"
const UPPERCASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALLCASES string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const SPECIALS string = "~!@#$%^&*()-_=+{}[]:;/?<>,.|\\"
const dictPath string = "usr/share/dict/words"

//
// Functions

// Up to you to decide which functions you want to add

////////////////////////////////////////////////////////////////////////////////
//
// Function     : generatePasword
// Description  : This is the function to generate the password.
//
// Inputs       : length - length of password
//                pattern - pattern of the file ("" if no pattern)
//                webflag - is this a web password?
// Outputs      : 0 if successful test, -1 if failure

func generatePasword(length int8, pattern string, webflag bool) string {

	pwd := "" // Start with nothing and add code
	var i int8 = 0

	fmt.Printf("length is %d\n", length)

	if length == 0 {
		if pattern == "" {
			length = 16
			fmt.Println("length changed to 16")
		} else {
			length = int8(len(pattern))
		}
	}

	if pattern == "" {
		for i = 0; i < length; i++ {
			pattern = pattern + "x"
		}
	}

	for i = 0; i < length; i++ {
		switch string(pattern[i]) {
		case "x":
			chartype := rand.Intn(3)
			if chartype == 0 {
				pwd += string(ALLCASES[rand.Intn(52)])
			} else if chartype == 1 {
				pwd += strconv.Itoa(rand.Intn(10))
			} else {
				pwd += string(SPECIALS[rand.Intn(29)])
			}
		case "d":
			pwd += strconv.Itoa(rand.Intn(10))
		case "c":
			pwd += string(ALLCASES[rand.Intn(52)])
		case "l":
			pwd += string(LOWERCASE[rand.Intn(26)])
		case "u":
			pwd += string(UPPERCASE[rand.Intn(26)])
		case "w":
			dictFile = os.Open(dictPath)

			if length > i+2 {
				if l, err := strconv.ParseInt(pattern[i+1:i+2], 10, 8); err == nil {
					// aqcuire word of length l
				} else if l, err = strconv.ParseInt(string(pattern[i+1]), 10, 8); err == nil {
					// acquire word of length l
				} else {
					// acquire any word
				}
			} else if length > i+1 {
				if l, err := strconv.ParseInt(string(pattern[i+1]), 10, 8); err == nil {
					// acquire word of length l
				} else {
					// acquire any word
				}
			} else {
				//acquire any word
			}

			dictFile.Close()
		case "s":
			pwd += string(SPECIALS[rand.Intn(29)])
		default:
			fmt.Println("Invalid character in given pattern.");
			return "NOPASSWORDCREATED"	// something went wrong
		}
	}

	// Now return the password
	return pwd
}

////////////////////////////////////////////////////////////////////////////////
//
// Function     : main
// Description  : The main function for the password generator program
//
// Inputs       : none
// Outputs      : 0 if successful test, -1 if failure

func main() {

	// Setup options for the program content
	rand.Seed(time.Now().UTC().UnixNano())
	helpflag := getopt.Bool('h', "", "help (this menu)")
	webflag := getopt.Bool('w', "", "web flag (no symbol characters, e.g., no &*...)")
	length := getopt.String('l', "", "length of password (in characters)")
	pattern := getopt.String('p', "", patternval)

	// Now parse the command line arguments
	err := getopt.Getopt(nil)
	if err != nil {
		// Handle error
		fmt.Fprintln(os.Stderr, err)
		getopt.Usage()
		os.Exit(-1)
	}

	// Get the flags
	fmt.Printf("helpflag [%t]\n", *helpflag)
	fmt.Printf("webflag [%t]\n", *webflag)
	fmt.Printf("length [%s]\n", *length)
	fmt.Printf("pattern [%s]\n", *pattern)
	// Normally, we we use getopt.Arg{#) to get the non-flag paramters

	// Safety check length parameter
	var plength int8 = 0
	if *length != "" {
		if _, err := strconv.Atoi(*length); err != nil {
			fmt.Printf("Bad length passed in [%s]\n", *length)
			fmt.Fprintln(os.Stderr, err)
			getopt.Usage()
			os.Exit(-1)
		}
		pl, _ := strconv.Atoi(*length)
		plength = int8(pl)
		if plength <= 0 || plength > 64 {
			plength = 16
		}
	}


	// Now generate the password and print it out
	pwd := generatePasword(plength, *pattern, *webflag)
	fmt.Printf("Generated password:  %s\n", pwd)

	// Return (no return code)
	return
}
