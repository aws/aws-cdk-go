package awsmediaconnect


// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &CfnGatewayProps{
//   	EgressCidrBlocks: []*string{
//   		jsii.String("egressCidrBlocks"),
//   	},
//   	Name: jsii.String("name"),
//   	Networks: []interface{}{
//   		&GatewayNetworkProperty{
//   			CidrBlock: jsii.String("cidrBlock"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html
//
type CfnGatewayProps struct {
	// The range of IP addresses that are allowed to contribute content or initiate output requests for flows communicating with this gateway.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-egresscidrblocks
	//
	EgressCidrBlocks *[]*string `field:"required" json:"egressCidrBlocks" yaml:"egressCidrBlocks"`
	// The name of the gateway.
	//
	// This name can not be modified after the gateway is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of networks in the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-networks
	//
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
}

