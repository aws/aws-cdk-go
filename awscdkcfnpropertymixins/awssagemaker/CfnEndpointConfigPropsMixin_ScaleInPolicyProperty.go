package awssagemaker


// Specifies how the endpoint releases instances when managed instance scaling scales in.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scaleInPolicyProperty := &ScaleInPolicyProperty{
//   	CooldownInMinutes: jsii.Number(123),
//   	MaximumStepSize: jsii.Number(123),
//   	Strategy: jsii.String("strategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-scaleinpolicy.html
//
type CfnEndpointConfigPropsMixin_ScaleInPolicyProperty struct {
	// The cooldown period, in minutes, after the last endpoint operation before the endpoint evaluates consolidation scale-in opportunities.
	//
	// Valid values are 5 to 1440. The default is 20.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-scaleinpolicy.html#cfn-sagemaker-endpointconfig-scaleinpolicy-cooldowninminutes
	//
	CooldownInMinutes *float64 `field:"optional" json:"cooldownInMinutes" yaml:"cooldownInMinutes"`
	// The maximum number of instances that the endpoint can terminate at a time during a consolidation scale-in operation.
	//
	// Valid values are 1 to 100. The default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-scaleinpolicy.html#cfn-sagemaker-endpointconfig-scaleinpolicy-maximumstepsize
	//
	MaximumStepSize *float64 `field:"optional" json:"maximumStepSize" yaml:"maximumStepSize"`
	// The strategy for scaling in instances.
	//
	// IDLE_RELEASE releases instances that have no hosted inference component copies. CONSOLIDATION consolidates inference component copies onto fewer instances to release more instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-scaleinpolicy.html#cfn-sagemaker-endpointconfig-scaleinpolicy-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

