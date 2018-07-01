package catalpha

import (
	"strings"
)

/*
	とりあえず、適当に単語を拾って挨拶する機能を作る
*/

func simpleResponse(in string) (out string) {
	for _, resp := range responses {
		if strings.Contains(in, resp.in) {
			return resp.out
		}
	}
	return ""
}

type response struct {
	in  string
	out string
}

var responses = []response{
	{"おはよ", "おはようにゃ"},
	{"おやすみ", "おやすみにゃ"},
	{"来た", "遅かったにゃ"},
	{"いってきま", "またにゃー"},
	{"行ってきま", "またきてにゃん"},
	{"調子どう？", "楽しいにゃ！"},
	{"ありがと", "ありがとにゃ！"},
	{"がんば", "私のためにがんばって:heart:"},
	{"楽し", "にゃ:musical_not:"},
	{"ねます", ":zzz:"},
	{"寝ます", "おやすみにゃ！"},
	{"お疲れ", "疲れたにゃ"},
	{"しまった", "気にしにゃーい"},
	{"乙です", "おつかれー"},
	{"おつで", "おつかれさまにゃ！"},
	{"すごい", "すごいにゃ！"},
	{"笑", "楽しいにゃん！"},
	{"好き", ":heart_decoration:"},
	{"かわいい", "通報するにゃ！"},
	{"猫", "にゃーんheart_eyes_cat"},
	{"ガチャ", ":money_with_wings:"},
}
