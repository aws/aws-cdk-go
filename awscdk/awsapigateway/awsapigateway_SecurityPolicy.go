package awsapigateway


// The minimum version of the SSL protocol that you want API Gateway to use for HTTPS connections.
//
// Example:
//   var acmCertificateForExampleCom interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("custom-domain"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acmCertificateForExampleCom,
//   	endpointType: apigateway.endpointType_EDGE,
//   	 // default is REGIONAL
//   	securityPolicy: apigateway.securityPolicy_TLS_1_2,
//   })
//
type SecurityPolicy string

const (
	// Cipher suite TLS 1.0.
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	// Cipher suite TLS 1.2.
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
)

