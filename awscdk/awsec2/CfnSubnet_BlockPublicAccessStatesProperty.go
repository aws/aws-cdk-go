package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockPublicAccessStatesProperty := &BlockPublicAccessStatesProperty{
//   	InternetGatewayBlockMode: jsii.String("internetGatewayBlockMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-subnet-blockpublicaccessstates.html
//
type CfnSubnet_BlockPublicAccessStatesProperty struct {
	// The mode of VPC BPA.
	//
	// Options here are off, block-bidirectional, block-ingress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-subnet-blockpublicaccessstates.html#cfn-ec2-subnet-blockpublicaccessstates-internetgatewayblockmode
	//
	InternetGatewayBlockMode *string `field:"optional" json:"internetGatewayBlockMode" yaml:"internetGatewayBlockMode"`
}

