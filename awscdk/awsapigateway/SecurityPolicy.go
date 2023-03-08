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
)

