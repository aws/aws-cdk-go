package awsroute53


// The record type.
type RecordType string

const (
	// route traffic to a resource, such as a web server, using an IPv4 address in dotted decimal notation.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#AFormat
	//
	RecordType_A RecordType = "A"
	// route traffic to a resource, such as a web server, using an IPv6 address in colon-separated hexadecimal format.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#AAAAFormat
	//
	RecordType_AAAA RecordType = "AAAA"
	// A CAA record specifies which certificate authorities (CAs) are allowed to issue certificates for a domain or subdomain.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#CAAFormat
	//
	RecordType_CAA RecordType = "CAA"
	// A CNAME record maps DNS queries for the name of the current record, such as acme.example.com, to another domain (example.com or example.net) or subdomain (acme.example.com or zenith.example.org).
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#CNAMEFormat
	//
	RecordType_CNAME RecordType = "CNAME"
	// A delegation signer (DS) record refers a zone key for a delegated subdomain zone.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#DSFormat
	//
	RecordType_DS RecordType = "DS"
	// An HTTPS resource record is a form of the Service Binding (SVCB) DNS record that provides extended configuration information, enabling a client to easily and securely connect to a service with an HTTP protocol.
	//
	// The configuration information is provided in parameters that allow the connection in one DNS query, rather than necessitating multiple DNS queries.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#HTTPSFormat
	//
	RecordType_HTTPS RecordType = "HTTPS"
	// An MX record specifies the names of your mail servers and, if you have two or more mail servers, the priority order.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#MXFormat
	//
	RecordType_MX RecordType = "MX"
	// A Name Authority Pointer (NAPTR) is a type of record that is used by Dynamic Delegation Discovery System (DDDS) applications to convert one value to another or to replace one value with another.
	//
	// For example, one common use is to convert phone numbers into SIP URIs.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#NAPTRFormat
	//
	RecordType_NAPTR RecordType = "NAPTR"
	// An NS record identifies the name servers for the hosted zone.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#NSFormat
	//
	RecordType_NS RecordType = "NS"
	// A PTR record maps an IP address to the corresponding domain name.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#PTRFormat
	//
	RecordType_PTR RecordType = "PTR"
	// A start of authority (SOA) record provides information about a domain and the corresponding Amazon Route 53 hosted zone.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#SOAFormat
	//
	RecordType_SOA RecordType = "SOA"
	// SPF records were formerly used to verify the identity of the sender of email messages.
	//
	// Instead of an SPF record, we recommend that you create a TXT record that contains the applicable value.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#SPFFormat
	//
	RecordType_SPF RecordType = "SPF"
	// An SRV record Value element consists of four space-separated values.
	//
	// The first three values are
	// decimal numbers representing priority, weight, and port. The fourth value is a domain name.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#SRVFormat
	//
	RecordType_SRV RecordType = "SRV"
	// A Secure Shell fingerprint record (SSHFP) identifies SSH keys associated with the domain name.
	//
	// SSHFP records must be secured with DNSSEC for a chain of trust to be established.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#SSHFPFormat
	//
	RecordType_SSHFP RecordType = "SSHFP"
	// You use an SVCB record to deliver configuration information for accessing service endpoints.
	//
	// The SVCB is a generic DNS record and can be used to negotiate parameters for a variety of application protocols.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#SVCBFormat
	//
	RecordType_SVCB RecordType = "SVCB"
	// You use a TLSA record to use DNS-Based Authentication of Named Entities (DANE).
	//
	// A TLSA record associates a certificate/public key with a Transport Layer Security (TLS) endpoint, and clients can validate the certificate/public key using a TLSA record signed with DNSSEC.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#TLSAFormat
	//
	RecordType_TLSA RecordType = "TLSA"
	// A TXT record contains one or more strings that are enclosed in double quotation marks (").
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html#TXTFormat
	//
	RecordType_TXT RecordType = "TXT"
)

