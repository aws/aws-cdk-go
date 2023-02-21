package awscertificatemanager


// Method used to assert ownership of the domain.
type ValidationMethod string

const (
	// Send email to a number of email addresses associated with the domain.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html
	//
	ValidationMethod_EMAIL ValidationMethod = "EMAIL"
	// Validate ownership by adding appropriate DNS records.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html
	//
	ValidationMethod_DNS ValidationMethod = "DNS"
)

