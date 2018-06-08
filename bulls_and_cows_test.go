package go_common

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type M struct {
	ImgUrl string
}

type D struct {
	Code    int
	Message string
	Data    M
}

func TestGuessAllCorrect(t *testing.T) {
	t.Run("guess4BullsAndNoCows", func(t *testing.T) {
		ret := getHint("1111", "1111")
		assert.Equal(t, "4A0B", ret, "guess4BullsAndNoCows failed")
	})

	t.Run("guessNoBullsAndCows", func(t *testing.T) {
		ret := getHint("1111", "2222")
		assert.Equal(t, "0A0B", ret, "guessNoBullsAndCows failed")
	})

	t.Run("guessOneBullAndNoCow", func(t *testing.T) {
		ret := getHint("1111", "1222")
		assert.Equal(t, "1A0B", ret, "guessOneBullAndNoCow failed")
	})

	t.Run("guessOneBullAndOneCow", func(t *testing.T) {
		ret := getHint("1112", "1323")
		assert.Equal(t, "1A1B", ret, "guessOneBullAndOneCow failed")
	})

	t.Run("guessOneBullAndOneCowWithDuplicateGuessNumber", func(t *testing.T) {
		ret := getHint("1112", "1223")
		assert.Equal(t, "1A1B", ret, "guessOneBullAndOneCowWithDuplicateGuessNumber failed")
	})

	t.Run("guessNoBullAndOneCowWithDuplicateGuessNumber", func(t *testing.T) {
		ret := getHint("1234", "0111")
		assert.Equal(t, "0A1B", ret, "guessNoBullAndOneCowWithDuplicateGuessNumber failed")
	})

	t.Run("guessNoBullAndFourCowWithDuplicateDigital", func(t *testing.T) {
		ret := getHint("1122", "2211")
		assert.Equal(t, "0A4B", ret, "guessNoBullAndFourCowWithDuplicateDigital failed")
	})

	t.Run("guessNoBullAndOneCowWithDuplicatePosition", func(t *testing.T) {
		ret := getHint("1122", "0001")
		assert.Equal(t, "0A1B", ret, "guessNoBullAndOneCowWithDuplicatePosition failed")
	})

	t.Run("test http get", func(t *testing.T) {
		res, err := http.Get("http://kong-http.api-zq-dev.baidao.com/aniu-service/v1/api/diagnosis/getImage?user=%E8%90%A7%E8%B6%85%E6%9D%B0&ticker=sh600000")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer res.Body.Close()
			contents, err := ioutil.ReadAll(res.Body)
			if err == nil {
				fmt.Println(string(contents))
				d := &D{}
				err = json.Unmarshal(contents, d)
				if err == nil {
					fmt.Printf("%v", d)
					res, err = http.Get(d.Data.ImgUrl)
					if err != nil {
						fmt.Printf("%v", err)
					} else {
						contents, err = ioutil.ReadAll(res.Body)
						if err != nil{
							fmt.Printf("%v", err)
						}else{
							fmt.Println(string(contents))
						}
					}
				} else {
					fmt.Printf("%v", err)
				}
			} else {
				fmt.Printf("%v", err)
			}
		}

	})
}
