package docx

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type SDT struct {
	XMLName    xml.Name       `xml:"w:sdt,omitempty"`
	Properties *SDTProperties `xml:"w:sdtPr,omitempty"`
	Content    *SDTContent    `xml:"w:sdtContent,omitempty"`
}

// SDTProperties represents the <w:sdtPr> tag
type SDTProperties struct {
	XMLName       xml.Name        `xml:"w:sdtPr,omitempty"`
	RunProperties *RunProperties  `xml:"w:rPr,omitempty"`
	ID            string          `xml:"w:id,attr,omitempty"`
	Placeholder   *SDTPlaceholder `xml:"w:placeholder,omitempty"`
	Alias         string          `xml:"w:alias,attr,omitempty"`
	Tag           string          `xml:"w:tag,attr,omitempty"`
}

type SDTPlaceholder struct {
	XMLName xml.Name `xml:"w:placeholder,omitempty"`
	DocPart *DocPart `xml:"w:docPart,omitempty"`
}

type DocPart struct {
	XMLName xml.Name `xml:"w:docPart,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// SDTContent represents the <w:sdtContent> tag
type SDTContent struct {
	Children []interface{}
}

func (s *SDT) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	println("tu")
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
			case "sdtContent":
				var value SDTContent
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				s.Content = &value
			case "sdtPr":
				var value SDTProperties
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				s.Properties = &value
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

func (p *SDTProperties) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	println("tu2")
	children := make([]interface{}, 0, 64)
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if tt, ok := t.(xml.StartElement); ok {
			var elem interface{}
			switch tt.Name.Local {
			case "rPr":
				var value RunProperties
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.RunProperties = &value
				elem = &value
			case "id":
				var value string
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.ID = value
				elem = &value
			case "placeholder":
				var value SDTPlaceholder
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Placeholder = &value
				elem = &value
			case "alias":
				var value string
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Alias = value
				elem = &value
			case "tag":
				var value string
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				p.Tag = value
				elem = &value
			default:
				err = d.Skip() // skip unsupported tags
				if err != nil {
					return err
				}
				continue
			}
			children = append(children, elem)
		}
	}
	return nil
}

func (p *SDTContent) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	println("tu2")
	children := make([]interface{}, 0, 64)
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if tt, ok := t.(xml.StartElement); ok {
			var elem interface{}
			switch tt.Name.Local {
			case "r":
				var value Run
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				elem = &value
			case "sdt":
				var value SDT
				err = d.DecodeElement(&value, &tt)
				prettyJSON, err := json.MarshalIndent(tt, "", "  ")
				if err != nil {
					fmt.Println("Failed to generate JSON:", err)
				}
				fmt.Println(string(prettyJSON))
				if err != nil {
					println(err.Error())
				}
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				elem = &value
			default:
				err = d.Skip() // skip unsupported tags
				if err != nil {
					return err
				}
				continue
			}
			children = append(children, elem)
		}
	}
	p.Children = children
	return nil
}
