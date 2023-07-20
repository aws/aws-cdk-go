package awsroute53


// *Private hosted zones only:* A complex type that contains information about an Amazon VPC.
//
// Route 53 Resolver uses the records in the private hosted zone to route traffic in that VPC.
//
// > For public hosted zones, omit `VPCs` , `VPCId` , and `VPCRegion` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCProperty := &VPCProperty{
//   	VpcId: jsii.String("vpcId"),
//   	VpcRegion: jsii.String("vpcRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html
//
type CfnHostedZone_VPCProperty struct {
	// *Private hosted zones only:* The ID of an Amazon VPC.
	//
	// > For public hosted zones, omit `VPCs` , `VPCId` , and `VPCRegion` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// *Private hosted zones only:* The region that an Amazon VPC was created in.
	//
	// > For public hosted zones, omit `VPCs` , `VPCId` , and `VPCRegion` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcregion
	//
	VpcRegion *string `field:"required" json:"vpcRegion" yaml:"vpcRegion"`
}

