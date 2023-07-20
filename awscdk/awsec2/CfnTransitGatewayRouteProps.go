package awsec2


// Properties for defining a `CfnTransitGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRouteProps := &CfnTransitGatewayRouteProps{
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//
//   	// the properties below are optional
//   	Blackhole: jsii.Boolean(false),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html
//
type CfnTransitGatewayRouteProps struct {
	// The ID of the transit gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-transitgatewayroutetableid
	//
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
	// Indicates whether to drop traffic that matches this route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-blackhole
	//
	Blackhole interface{} `field:"optional" json:"blackhole" yaml:"blackhole"`
	// The CIDR block used for destination matches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

