package awseks


// The node auto repair configuration for the node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeRepairConfigProperty := &NodeRepairConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html
//
type CfnNodegroup_NodeRepairConfigProperty struct {
	// Specifies whether to enable node auto repair for the node group.
	//
	// Node auto repair is disabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

