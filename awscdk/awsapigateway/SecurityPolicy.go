package awsapigateway


// The minimum version of the SSL protocol that you want API Gateway to use for HTTPS connections.
//
// Example:
//   var acmCertificateForExampleCom interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("custom-domain"), &DomainNameProps{
//   	DomainName: jsii.String("example.com"),
//   	Certificate: acmCertificateForExampleCom,
//   	EndpointType: apigateway.EndpointType_EDGE,
//   	 // default is REGIONAL
//   	SecurityPolicy: apigateway.SecurityPolicy_TLS_1_2,
//   })
//
type SecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
	// Cipher suite TLS 1.3 for regional/private endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS13_1_3_2025_09 SecurityPolicy = "TLS13_1_3_2025_09"
	// Cipher suite TLS 1.3 (FIPS compliant) for regional/private endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS13_1_3_FIPS_2025_09 SecurityPolicy = "TLS13_1_3_FIPS_2025_09"
	// Cipher suite TLS 1.3 and TLS 1.2 with post-quantum cryptography for regional/private endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS13_1_2_PQ_2025_09 SecurityPolicy = "TLS13_1_2_PQ_2025_09"
	// Cipher suite TLS 1.3 and TLS 1.2 with Perfect Forward Secrecy and post-quantum cryptography for regional/private endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 SecurityPolicy = "TLS13_1_2_PFS_PQ_2025_09"
	// Cipher suite TLS 1.3 for edge-optimized endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS13_2025_EDGE SecurityPolicy = "TLS13_2025_EDGE"
	// Cipher suite TLS 1.2 with Perfect Forward Secrecy for edge-optimized endpoints.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS12_PFS_2025_EDGE SecurityPolicy = "TLS12_PFS_2025_EDGE"
	// Cipher suite TLS 1.2 for edge-optimized endpoints (legacy).
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html
	//
	SecurityPolicy_TLS12_2018_EDGE SecurityPolicy = "TLS12_2018_EDGE"
)

