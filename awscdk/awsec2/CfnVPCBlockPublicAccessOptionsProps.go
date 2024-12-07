package awsec2


// Properties for defining a `CfnVPCBlockPublicAccessOptions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCBlockPublicAccessOptionsProps := &CfnVPCBlockPublicAccessOptionsProps{
//   	InternetGatewayBlockMode: jsii.String("internetGatewayBlockMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessoptions.html
//
type CfnVPCBlockPublicAccessOptionsProps struct {
	// The desired Block Public Access mode for Internet Gateways in your account.
	//
	// We do not allow to create in a off mode as this is the default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessoptions.html#cfn-ec2-vpcblockpublicaccessoptions-internetgatewayblockmode
	//
	InternetGatewayBlockMode *string `field:"required" json:"internetGatewayBlockMode" yaml:"internetGatewayBlockMode"`
}

