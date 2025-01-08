package docx

import (
	"io/fs"
)

// UseTemplate will replace template files
func (f *Docx) UseTemplate(template string, tmpfslst []string, tmplfs fs.FS) *Docx {
	f.template = template
	f.tmplfs = tmplfs
	f.tmpfslst = tmpfslst
	return f
}

// WithDefaultTheme use default theme embeded
func (f *Docx) WithDefaultTheme() *Docx {
	return f.UseTemplate("default", A4TemplateFilesList, TemplateXMLFS)
}

// WithA3Page use A3 PageSize
func (f *Docx) WithA3Page() *Docx {
	sectpr := &SectPr{
		PgSz: &PgSz{
			W: 16838,
			H: 23811,
		},
	}
	f.Document.Body.Items = append(f.Document.Body.Items, sectpr)
	return f
}

// WithA4Page use A4 PageSize
func (f *Docx) WithA4Page() *Docx {
	sectpr := &SectPr{
		PgSz: &PgSz{
			W: 11906,
			H: 16838,
		},
	}
	f.Document.Body.Items = append(f.Document.Body.Items, sectpr)
	return f
}
