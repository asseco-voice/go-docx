package docx

import (
	"encoding/xml"
	"io"
	"strings"
)

/*
<w:r w:rsidR="007427C7"><w:fldChar w:fldCharType="begin"> <w:ffData> <w:name w:val="Pareiskejas"/> <w:enabled/> <w:calcOnExit w:val="0"/> <w:textInput><w:default w:val="UAB „Vilniaus vandenys“."/></w:textInput></w:ffData></w:fldChar></w:r>
*/

type Default struct {
	XMLName xml.Name `xml:"w:default,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

type TextInput struct {
	XMLName xml.Name `xml:"w:textInput,omitempty"`
	Default *Default `xml:"w:default,omitempty"`
}

type CalcOnExit struct {
	XMLName xml.Name `xml:"w:calcOnExit,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

type Enabled struct {
	XMLName xml.Name `xml:"w:enabled,omitempty"`
}

type Name struct {
	XMLName xml.Name `xml:"w:name,omitempty"`
	Val     string   `xml:"w:val,attr,omitempty"`
}

type FfData struct {
	XMLName    xml.Name    `xml:"w:ffData,omitempty"`
	Name       *Name       `xml:"w:name,omitempty"`
	Enabled    *Enabled    `xml:"w:enabled,omitempty"`
	CalcOnExit *CalcOnExit `xml:"w:calcOnExit,omitempty"`
	TextInput  *TextInput  `xml:"w:textInput,omitempty"`
}

type FldChar struct {
	XMLName     xml.Name `xml:"w:fldChar,omitempty"`
	FldCharType string   `xml:"w:fldCharType,attr,omitempty"`
	FfData      *FfData  `xml:"w:ffData,omitempty"`
}

func (p *FldChar) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			case "ffData":
				var value FfData
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.FfData = &value
			default:
				err = d.Skip()
				if err != nil {
					return err
				}
				continue
			}
		}
	}
	return nil
}

func (p *FfData) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			case "name":
				var value Name
				err = d.DecodeElement(&value, &tt)
				v := getAtt(tt.Attr, "val")
				if v == "" {
					continue
				}
				value.Val = v
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Name = &value
			case "enabled":
				var value Enabled
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Enabled = &value
			case "calcOnExit":
				var value CalcOnExit
				err = d.DecodeElement(&value, &tt)
				v := getAtt(tt.Attr, "val")
				if v == "" {
					continue
				}
				value.Val = v
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.CalcOnExit = &value
			case "textInput":
				var value TextInput
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.TextInput = &value
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
func (p *TextInput) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			case "default":
				var value Default
				err = d.DecodeElement(&value, &tt)
				v := getAtt(tt.Attr, "val")
				if v == "" {
					continue
				}
				value.Val = v
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Default = &value
			default:
				err = d.Skip()
				if err != nil {
					return err
				}
				continue
			}
		}
	}
	return nil
}
