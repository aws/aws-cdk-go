package awsapigateway


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
type EndpointType string

const (
	// For an edge-optimized API and its custom domain name.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional API and its custom domain name.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	// For a private API and its custom domain name.
	EndpointType_PRIVATE EndpointType = "PRIVATE"
)

