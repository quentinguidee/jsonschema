package jsonschema

type Format string

const (
	FormatDateTime Format = "date-time"
	FormatTime     Format = "time"
	FormatDate     Format = "date"
	FormatDuration Format = "duration"

	FormatEmail    Format = "email"
	FormatIdnEmail Format = "idn-email"

	FormatHostname    Format = "hostname"
	FormatIdnHostname Format = "idn-hostname"

	FormatIpv4 Format = "ipv4"
	FormatIpv6 Format = "ipv6"

	FormatUUID         Format = "uuid"
	FormatURI          Format = "uri"
	FormatURIReference Format = "uri-reference"
	FormatIRI          Format = "iri"
	FormatIRIReference Format = "iri-reference"

	FormatURITemplate Format = "uri-template"

	FormatJsonPointer         Format = "json-pointer"
	FormatRelativeJsonPointer Format = "relative-json-pointer"

	FormatRegex Format = "regex"
)
