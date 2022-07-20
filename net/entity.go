package net

import (
	"bytes"
)

const (
	ContentType string = "Content-Type: text/plain"
	LineFeed    string = "\r\n"
	CHARSET     string = "UTF-8"
	BOUNDARY    string = "e2a540ab4e6c5ed79c01157c255a2b5007e157d7"
	BoundaryFix string = "--"
	UserAgent   string = "This Library Created by Bojja Vighneswar Rao, https://vighnesh.org"
)

type Entity struct {
	bytes.Buffer
}

func (me *Entity) AddHeader(name, value string) {
	me.WriteString(name)
	me.WriteString(":")
	me.WriteString(value)
	me.WriteString(LineFeed)
}

func (me *Entity) AddTextBody(typ, value string) {
	me.WriteString(BoundaryFix)
	me.WriteString(BOUNDARY)
	me.WriteString(LineFeed)
	me.WriteString("Content-Disposition: form-data; name=\"" + typ + "\"")
	me.WriteString(LineFeed)
	me.WriteString(ContentType)
	me.WriteString("; charset=")
	me.WriteString(CHARSET)
	me.WriteString(LineFeed)
	me.WriteString(LineFeed)
	me.WriteString(value)
	me.WriteString(LineFeed)
}

func (me *Entity) Build() []byte {
	me.WriteString(LineFeed)
	me.WriteString(BoundaryFix)
	me.WriteString(BOUNDARY)
	me.WriteString(BoundaryFix)
	me.WriteString(LineFeed)

	return me.Bytes()
}
