package docx

import (
	"encoding/xml"
	"io"
	"strconv"
)

// NumProperties <w:numPr>
type NumProperties struct {
	XMLName xml.Name `xml:"w:numPr,omitempty"`
	Ilvl    *Ilvl
	NumId   *NumId
	Lvl     *Lvl
}

// Ilvl ...
type Ilvl struct {
	XMLName xml.Name `xml:"w:ilvl,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// NumId ...
type NumId struct {
	XMLName xml.Name `xml:"w:numId,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

type KeepNext struct {
	XMLName xml.Name `xml:"w:keepNext,omitempty"`
	Val     int      `xml:"w:val,attr,omitempty"`
}

func (p *NumProperties) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if tt, ok := t.(xml.StartElement); ok {
			switch tt.Name.Local {
			case "ilvl":
				var value Ilvl
				v := getAtt(tt.Attr, "val")
				if v == "" {
					continue
				}
				value.Val = v
				p.Ilvl = &value
			case "numId":
				var value NumId
				v := getAtt(tt.Attr, "val")
				if v == "" {
					continue
				}
				value.Val = v
				p.NumId = &value
			default:
				err = d.Skip() // skip unsupported tags
				if err != nil {
					return err
				}
				continue
			}
		}
	}

	return nil
}

func (k *KeepNext) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			k.Val, err = strconv.Atoi(attr.Value)
			if err != nil {
				return
			}
		default:
			// ignore other attributes
		}
	}
	// Consume the end element
	_, err = d.Token()
	return
}
