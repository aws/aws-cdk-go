package awsmediaconnect


// Specifies the configuration settings for individual NDI discovery servers.
//
// A maximum of 3 servers is allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ndiDiscoveryServerConfigProperty := &NdiDiscoveryServerConfigProperty{
//   	DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   	VpcInterfaceAdapter: jsii.String("vpcInterfaceAdapter"),
//
//   	// the properties below are optional
//   	DiscoveryServerPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndidiscoveryserverconfig.html
//
type CfnFlow_NdiDiscoveryServerConfigProperty struct {
	// The unique network address of the NDI discovery server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndidiscoveryserverconfig.html#cfn-mediaconnect-flow-ndidiscoveryserverconfig-discoveryserveraddress
	//
	DiscoveryServerAddress *string `field:"required" json:"discoveryServerAddress" yaml:"discoveryServerAddress"`
	// The identifier for the Virtual Private Cloud (VPC) network interface used by the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndidiscoveryserverconfig.html#cfn-mediaconnect-flow-ndidiscoveryserverconfig-vpcinterfaceadapter
	//
	VpcInterfaceAdapter *string `field:"required" json:"vpcInterfaceAdapter" yaml:"vpcInterfaceAdapter"`
	// The port for the NDI discovery server.
	//
	// Defaults to 5959 if a custom port isn't specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndidiscoveryserverconfig.html#cfn-mediaconnect-flow-ndidiscoveryserverconfig-discoveryserverport
	//
	DiscoveryServerPort *float64 `field:"optional" json:"discoveryServerPort" yaml:"discoveryServerPort"`
}

