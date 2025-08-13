package awspcs


// An endpoint available for interaction with the scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &EndpointProperty{
//   	Port: jsii.String("port"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Ipv6Address: jsii.String("ipv6Address"),
//   	PublicIpAddress: jsii.String("publicIpAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html
//
type CfnCluster_EndpointProperty struct {
	// The endpoint's connection port number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-port
	//
	Port *string `field:"required" json:"port" yaml:"port"`
	// The endpoint's private IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-privateipaddress
	//
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Indicates the type of endpoint running at the specific IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The endpoint's IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-ipv6address
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The endpoint's public IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-endpoint.html#cfn-pcs-cluster-endpoint-publicipaddress
	//
	PublicIpAddress *string `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
}

