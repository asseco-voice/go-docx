package docx

import (
	"encoding/xml"
	"io"
)

type Numbering struct {
	XMLName xml.Name `xml:"w:numbering"`
	XMLW    string   `xml:"xmlns:w,attr"`             // cannot be unmarshalled in
	XMLR    string   `xml:"xmlns:r,attr,omitempty"`   // cannot be unmarshalled in
	XMLWP   string   `xml:"xmlns:wp,attr,omitempty"`  // cannot be unmarshalled in
	XMLWPS  string   `xml:"xmlns:wps,attr,omitempty"` // cannot be unmarshalled in
	XMLWPC  string   `xml:"xmlns:wpc,attr,omitempty"` // cannot be unmarshalled in
	XMLWPG  string   `xml:"xmlns:wpg,attr,omitempty"` // cannot be unmarshalled in

	AbstractNums *[]AbstractNum
	Nums         *[]Num `xml:"num,omitempty"`

	file *Docx
}

type Num struct {
	XMLName        xml.Name `xml:"num,omitempty"`
	NumID          string   `xml:"w:numId,attr"`
	*AbstractNumID `xml:"w:abstractNumId,omitempty"`
	// AbstractNumID *AbstractNumID `xml:"w:abstractNumId",omitempty`
}

type AbstractNumID struct {
	XMLName xml.Name `xml:"abstractNumId,omitempty"`
	// CommonAttrVal *CommonAttrVal
	Val string `xml:"w:val,attr"`
}

type AbstractNum struct {
	XMLName        xml.Name `xml:"w:abstractNum,omitempty"`
	AbstractNumId  string   `xml:"w:abstractNumId"`
	Lvl            *[]Lvl   `xml:"lvl"`
	Nsid           *Nsid
	MultiLevelType *MultiLevelType
	Tmpl           *Tmpl
}

type Nsid struct {
	XMLName xml.Name `xml:"w:nsid"`
	Val     string   `xml:"w:val"`
}

type MultiLevelType struct {
	XMLName xml.Name `xml:"w:multiLevelType"`
	Val     string   `xml:"w:val"`
}

type Tmpl struct {
	XMLName xml.Name `xml:"w:tmpl"`
	Val     string   `xml:"w:val"`
}

type Lvl struct {
	XMLName   xml.Name `xml:"w:lvl"`
	Ilvl      string   `xml:"w:ilvl"`
	Tplc      string   `xml:"w:tplc,attr"`
	Tentative string   `xml:"w:tentative"`
	Start     *Start
	NumFmt    *NumFmt
	LvlText   *LvlText
	LvlJc     *LvlJc
	PPr       *[]ParagraphProperties
	RPr       *RunProperties
}

type Start struct {
	XMLName xml.Name `xml:"w:start"`
	Val     string   `xml:"w:val"`
}

type NumFmt struct {
	XMLName xml.Name `xml:"w:numFmt"`
	Val     string   `xml:"w:val"`
}

type LvlText struct {
	XMLName xml.Name `xml:"w:lvlText"`
	Val     string   `xml:"w:val"`
}

type LvlJc struct {
	XMLName xml.Name `xml:"w:lvlJc"`
	Val     string   `xml:"w:val"`
}

func (n *Numbering) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			case "abstractNum":
				var an AbstractNum
				err = d.DecodeElement(&an, &tt)
				if err != nil {
					return err
				}
				if n.AbstractNums == nil {
					n.AbstractNums = &[]AbstractNum{}
				}
				*n.AbstractNums = append(*n.AbstractNums, an)
			case "num":
				var num Num
				err = d.DecodeElement(&num, &tt)
				if err != nil {
					return err
				}
				if n.Nums == nil {
					n.Nums = &[]Num{}
				}
				*n.Nums = append(*n.Nums, num)
			default:
				// ignore other attributes
			}
		}
	}

	return nil
}

// AbstructNum UnmarshalXML ...
func (a *AbstractNum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "abstractNumId":
			a.AbstractNumId = attr.Value
		default:
			// ignore other attributes
		}
	}

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
			case "lvl":
				l := Lvl{}
				err = d.DecodeElement(&l, &tt)
				if a.Lvl == nil {
					a.Lvl = &[]Lvl{}
				}
				*a.Lvl = append(*a.Lvl, l)
			case "nsid":
				n := NewNSID()
				err = d.DecodeElement(n, &tt)
				(*a).Nsid = n
			case "multiLevelType":
				m := NewMultiLevelType()
				err = d.DecodeElement(m, &tt)
				(*a).MultiLevelType = m
			case "tmpl":
				t := NewTmpl()
				err = d.DecodeElement(t, &tt)
				(*a).Tmpl = t
			}
		}
	}

	return
}

// Lvl UnmarshalXML ...
func (l *Lvl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "ilvl":
			l.Ilvl = attr.Value
		case "tplc":
			l.Tplc = attr.Value
		case "tentative":
			l.Tentative = attr.Value
		default:
			// ignore other attributes
		}
	}

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
			case "pPr":
				var p ParagraphProperties
				err = d.DecodeElement(&p, &tt)
				if err != nil {
					return err
				}
				if l.PPr == nil {
					l.PPr = &[]ParagraphProperties{}
				}
				*l.PPr = append(*l.PPr, p)
			case "start":
				s := NewStart()
				err = d.DecodeElement(s, &tt)
				if err != nil {
					return err
				}
				l.Start = s
			case "numFmt":
				n := NewNumFmt()
				err = d.DecodeElement(n, &tt)
				if err != nil {
					return err
				}
				l.NumFmt = n
			case "lvlText":
				lt := NewLvlText()
				err = d.DecodeElement(lt, &tt)
				if err != nil {
					return err
				}
				l.LvlText = lt
			case "lvlJc":
				lj := NewLvlJc()
				err = d.DecodeElement(lj, &tt)
				if err != nil {
					return err
				}
				l.LvlJc = lj
			default:
				// ignore other attributes
			}
		}
	}

	return
}
func (n *Num) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "numId":
			n.NumID = attr.Value
		default:
			// ignore other attributes
		}
	}
	// attr ではなく要素の値を取得する
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
			case "abstractNumId":
				an := NewAbstractNumID()
				err = d.DecodeElement(an, &tt)
				if err != nil {
					return err
				}
				n.AbstractNumID = an
			default:
				// ignore other attributes
			}
		}
	}
	return
}

func NewAbstractNumID() *AbstractNumID {
	return &AbstractNumID{Val: ""}
}
func NewNSID() *Nsid {
	return &Nsid{Val: ""}
}
func NewMultiLevelType() *MultiLevelType {
	return &MultiLevelType{Val: ""}
}
func NewTmpl() *Tmpl {
	return &Tmpl{Val: ""}
}
func NewStart() *Start {
	return &Start{Val: ""}
}
func NewNumFmt() *NumFmt {
	return &NumFmt{Val: ""}
}
func NewLvlText() *LvlText {
	return &LvlText{Val: ""}
}
func NewLvlJc() *LvlJc {
	return &LvlJc{Val: ""}
}
