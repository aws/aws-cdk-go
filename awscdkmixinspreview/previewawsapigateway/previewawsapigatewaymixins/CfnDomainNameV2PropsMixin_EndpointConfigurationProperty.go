package previewawsapigatewaymixins


// The endpoint configuration to indicate the types of endpoints an API (RestApi) or its custom domain name (DomainName) has and the IP address types that can invoke it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointConfigurationProperty := &EndpointConfigurationProperty{
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainnamev2-endpointconfiguration.html
//
type CfnDomainNameV2PropsMixin_EndpointConfigurationProperty struct {
	// The IP address types that can invoke an API (RestApi) or a DomainName.
	//
	// Use `ipv4` to allow only IPv4 addresses to invoke an API or DomainName, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke an API or a DomainName. For the `PRIVATE` endpoint type, only `dualstack` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainnamev2-endpointconfiguration.html#cfn-apigateway-domainnamev2-endpointconfiguration-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A list of endpoint types of an API (RestApi) or its custom domain name (DomainName).
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is `"EDGE"` . For a regional API and its custom domain name, the endpoint type is `REGIONAL` . For a private API, the endpoint type is `PRIVATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainnamev2-endpointconfiguration.html#cfn-apigateway-domainnamev2-endpointconfiguration-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

