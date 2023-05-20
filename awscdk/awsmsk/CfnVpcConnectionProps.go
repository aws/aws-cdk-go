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
	// The type of private link authentication.
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// The list of subnets in the client VPC to connect to.
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// The security groups to attach to the ENIs for the broker nodes.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The Amazon Resource Name (ARN) of the cluster.
	TargetClusterArn *string `field:"required" json:"targetClusterArn" yaml:"targetClusterArn"`
	// The VPC id of the remote client.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Create tags when creating the VPC connection.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

