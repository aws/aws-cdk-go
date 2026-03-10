package awsapigateway


// The endpoint access mode for the domain name.
//
// When using enhanced security policies (those starting with `SecurityPolicy_`),
// you must set the endpoint access mode to either `STRICT` or `BASIC`.
// Use `STRICT` for production workloads requiring the highest security.
// Use `BASIC` for migration scenarios or certain application architectures.
//
// Example:
//   var acmCertificateForExampleCom interface{}
//
//
//   // For regional or private APIs with enhanced security policy
//   // For regional or private APIs with enhanced security policy
//   apigateway.NewDomainName(this, jsii.String("custom-domain-tls13"), &DomainNameProps{
//   	DomainName: jsii.String("example.com"),
//   	Certificate: acmCertificateForExampleCom,
//   	SecurityPolicy: apigateway.SecurityPolicy_TLS13_1_3_2025_09,
//   	 // TLS 1.3
//   	EndpointAccessMode: apigateway.EndpointAccessMode_STRICT,
//   })
//
//   // For edge-optimized APIs with enhanced security policy
//   // For edge-optimized APIs with enhanced security policy
//   apigateway.NewDomainName(this, jsii.String("custom-domain-edge-tls13"), &DomainNameProps{
//   	DomainName: jsii.String("example.com"),
//   	Certificate: acmCertificateForExampleCom,
//   	EndpointType: apigateway.EndpointType_EDGE,
//   	SecurityPolicy: apigateway.SecurityPolicy_TLS13_2025_EDGE,
//   	 // Enhanced security policy for edge
//   	EndpointAccessMode: apigateway.EndpointAccessMode_STRICT,
//   })
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-security-policies.html#apigateway-security-policies-endpoint-access-mode
//
type EndpointAccessMode string

const (
	// Strict mode - only accepts connections from clients using the specified security policy.
	//
	// Recommended for production workloads.
	EndpointAccessMode_STRICT EndpointAccessMode = "STRICT"
	// Basic mode - one of the two valid endpoint access modes for enhanced security policies.
	//
	// Suitable for migration scenarios or certain application architectures.
	// Note: legacy security policies (TLS_1_0, TLS_1_2) do not use this attribute.
	EndpointAccessMode_BASIC EndpointAccessMode = "BASIC"
)

