package jsonschema

type ContentEncoding string

const (
	ContentEncodingQuotedPrintable ContentEncoding = "quoted-printable"
	ContentEncodingBase16          ContentEncoding = "base16"
	ContentEncodingBase32          ContentEncoding = "base32"
	ContentEncodingBase64          ContentEncoding = "base64"
)
