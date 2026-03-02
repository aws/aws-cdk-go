package previewawsbatchmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   computeScalingPolicyProperty := &ComputeScalingPolicyProperty{
//   	MinScaleDownDelayMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computescalingpolicy.html
//
type CfnComputeEnvironmentPropsMixin_ComputeScalingPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computescalingpolicy.html#cfn-batch-computeenvironment-computescalingpolicy-minscaledowndelayminutes
	//
	MinScaleDownDelayMinutes *float64 `field:"optional" json:"minScaleDownDelayMinutes" yaml:"minScaleDownDelayMinutes"`
}

