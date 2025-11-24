package mixinsawscassandra


// Amazon Keyspaces supports the `target tracking` auto scaling policy.
//
// With this policy, Amazon Keyspaces auto scaling ensures that the table's ratio of consumed to provisioned capacity stays at or near the target value that you specify. You define the target value as a percentage between 20 and 90.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingPolicyProperty := &ScalingPolicyProperty{
//   	TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   		DisableScaleIn: jsii.Boolean(false),
//   		ScaleInCooldown: jsii.Number(123),
//   		ScaleOutCooldown: jsii.Number(123),
//   		TargetValue: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-scalingpolicy.html
//
type CfnTablePropsMixin_ScalingPolicyProperty struct {
	// The auto scaling policy that scales a table based on the ratio of consumed to provisioned capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-scalingpolicy.html#cfn-cassandra-table-scalingpolicy-targettrackingscalingpolicyconfiguration
	//
	TargetTrackingScalingPolicyConfiguration interface{} `field:"optional" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

