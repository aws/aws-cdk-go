package mixinsawsec2


// Describes the target groups to attach to a Spot Fleet.
//
// Spot Fleet registers the running Spot Instances with these target groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetGroupsConfigProperty := &TargetGroupsConfigProperty{
//   	TargetGroups: []interface{}{
//   		&TargetGroupProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-targetgroupsconfig.html
//
type CfnSpotFleetPropsMixin_TargetGroupsConfigProperty struct {
	// One or more target groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-targetgroupsconfig.html#cfn-ec2-spotfleet-targetgroupsconfig-targetgroups
	//
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

