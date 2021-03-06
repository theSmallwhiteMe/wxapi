package tools

//xml map 转换

import (
	"encoding/xml"
	"io"
)

type XmlMap map[string]string

type XmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m XmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(XmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *XmlMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = XmlMap{}
	for {
		var e XmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}
