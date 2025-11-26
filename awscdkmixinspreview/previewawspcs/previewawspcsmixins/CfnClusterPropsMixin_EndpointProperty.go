package previewawspcsmixins


// An endpoint available for interaction with the scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointProperty := &EndpointProperty{
//   	Ipv6Address: jsii.String("ipv6Address"),
//   	Port: jsii.String("port"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PublicIpAddress: jsii.String("publicIpAddress"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html
//
type CfnClusterPropsMixin_EndpointProperty struct {
	// The endpoint's IPv6 address.
	//
	// Example: `2001:db8::1`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-ipv6address
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The endpoint's connection port number.
	//
	// Example: `1234`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
	// For clusters that use IPv4, this is the endpoint's private IP address.
	//
	// Example: `10.1.2.3`
	//
	// For clusters configured to use IPv6, this is an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The endpoint's public IP address.
	//
	// Example: `192.0.2.1`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-publicipaddress
	//
	PublicIpAddress *string `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
	// Indicates the type of endpoint running at the specific IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

