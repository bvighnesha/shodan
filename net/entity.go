package net

import (
	"bytes"
)

const (
	CONTENT_TYPE string = "Content-Type: text/plain"
	LINE_FEED    string = "\r\n"
	CHARSET      string = "UTF-8"
	BOUNDARY     string = "e2a540ab4e6c5ed79c01157c255a2b5007e157d7"
	BOUNDARY_FIX string = "--"
	USER_AGENT   string = "This Library Created by Bojja Vighneswar Rao, https://vighnesh.org"
)

type Entity struct {
	bytes.Buffer
}

func (me *Entity) AddHeader(name, value string) {

	me.WriteString(name)
	me.WriteString(":")
	me.WriteString(value)
	me.WriteString(LINE_FEED)
}

func (me *Entity) AddTextBody(typ, value string) {
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(BOUNDARY)
	me.WriteString(LINE_FEED)
	me.WriteString("Content-Disposition: form-data; name=\"" + typ + "\"")
	me.WriteString(LINE_FEED)
	me.WriteString(CONTENT_TYPE)
	me.WriteString("; charset=")
	me.WriteString(CHARSET)
	me.WriteString(LINE_FEED)
	me.WriteString(LINE_FEED)
	me.WriteString(value)
	me.WriteString(LINE_FEED)
}

func (me *Entity) Build() []byte {
	me.WriteString(LINE_FEED)
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(BOUNDARY)
	me.WriteString(BOUNDARY_FIX)
	me.WriteString(LINE_FEED)
	return me.Bytes()
}
