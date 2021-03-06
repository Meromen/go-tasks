package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tts2media "github.com/pqyptixa/tts2media"
)

// Language: language of the speaker, aka "voice name" in espeak; values: "af", "bs", "ca", "cs",
// "cy", "de", "en", "en-sc", "en-uk", "en-uk-north", "en-uk-rp", "en-uk-wmids", "en-us", "en-wi",
// "eo", "es", "es-la", "fi", "fr", "fr-be", "grc", "hr", "hu", "id", "is", "it", "jbo", "ku",
// "la", "lv", "mk", "nl", "no", "pl", "pt-pt", "pt", "ro", "ru", "sk", "sq", "sr", "sv", "sw",
// "tr", "vi", "zh", "zh-yue", "hi", "el", "ta"
//
// Speed: speed in words per minute; values: from "80" to "390"
//
// Gender: gender of the speaker; values: "m" for male, and "f" for female. note: not all voices support different genders
//
// Altvoice: alternative voice; values: from "0" to "5"
//
// Quality: quality of the output MP3/OGG audio; values: "high", "medium" or "low"
//
// Pitch: pitch adjustment; values: from "0" to "99"

func main() {
	text, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	espeak := &tts2media.EspeakSpeech{
		Text:     string(text), // text to turn to speech
		Lang:     "en",         // language
		Speed:    "190",        // speed
		Gender:   "f",          // gender
		Altvoice: "0",          // alternative voice
		Quality:  "high",       // quality of output mp3/ogg audio
		Pitch:    "50",         // pitch
	}

	media, err := espeak.NewEspeakSpeech()
	fmt.Println(media)
	err = media.ToAudio()
	if err != nil {
		fmt.Println(err)
	}
	os.Rename("./"+media.Filename+".mp3", "testAudio.mp3")
	os.Remove("./" + media.Filename + ".ogg")
	os.Remove("./" + media.Filename + ".mp3")

	media.RemoveWAV()
}
