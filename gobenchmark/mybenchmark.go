package mybenchmark

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"strings"
	"bytes"
)

type Foobar struct {
	Foo string
	Bar int
}

func MarshalFoobar(fb Foobar) error {

	_, err := json.MarshalIndent(fb, "", "	")
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func EncodeFoobar(stream string) error {

	var fb Foobar

	enc := json.NewEncoder(bytes.NewBufferString(stream))
	enc.SetIndent("", "	")

	return enc.Encode(&fb)
}

func UnmarshalFoobar(stream string) (Foobar, error) {

	var fb Foobar

	data, err := ioutil.ReadAll(strings.NewReader(stream))
	if err != nil {
		log.Fatal(err)
		return Foobar{}, err
	}

	err = json.Unmarshal(data, &fb)
	return fb, err
}

func DecodeFoobar(stream string) (Foobar, error) {
	
	dec := json.NewDecoder(strings.NewReader(stream))

	var fb Foobar

	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		err := dec.Decode(&fb)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return fb, err
}