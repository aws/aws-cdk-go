package awsmsk


// Properties for defining a `CfnVpcConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcConnectionProps := &CfnVpcConnectionProps{
//   	Authentication: jsii.String("authentication"),
//   	ClientSubnets: []*string{
//   		jsii.String("clientSubnets"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	TargetClusterArn: jsii.String("targetClusterArn"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html
//
type CfnVpcConnectionProps struct {
	// The type of private link authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-authentication
	//
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-clientsubnets
	//
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-securitygroups
	//
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The Amazon Resource Name (ARN) of the target cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-targetclusterarn
	//
	TargetClusterArn *string `field:"required" json:"targetClusterArn" yaml:"targetClusterArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// A key-value pair to associate with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-vpcconnection.html#cfn-msk-vpcconnection-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

