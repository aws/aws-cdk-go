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
type CfnVpcConnectionProps struct {
	// `AWS::MSK::VpcConnection.Authentication`.
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// `AWS::MSK::VpcConnection.ClientSubnets`.
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// `AWS::MSK::VpcConnection.SecurityGroups`.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// `AWS::MSK::VpcConnection.TargetClusterArn`.
	TargetClusterArn *string `field:"required" json:"targetClusterArn" yaml:"targetClusterArn"`
	// `AWS::MSK::VpcConnection.VpcId`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// `AWS::MSK::VpcConnection.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

