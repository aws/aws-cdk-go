package previewawsapigatewaymixins


// The `EndpointConfiguration` property type specifies the endpoint types and IP address types of an Amazon API Gateway domain name.
//
// `EndpointConfiguration` is a property of the [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html
//
type CfnDomainNamePropsMixin_EndpointConfigurationProperty struct {
	// The IP address types that can invoke this DomainName.
	//
	// Use `ipv4` to allow only IPv4 addresses to invoke this DomainName, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke this DomainName. For the `PRIVATE` endpoint type, only `dualstack` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A list of endpoint types of an API (RestApi) or its custom domain name (DomainName).
	//
	// For an edge-optimized API and its custom domain name, the endpoint type is `"EDGE"` . For a regional API and its custom domain name, the endpoint type is `REGIONAL` . For a private API, the endpoint type is `PRIVATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

