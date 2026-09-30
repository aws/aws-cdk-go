package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsSettingsProperty := &EcsSettingsProperty{
//   	ContainerInsights: jsii.String("containerInsights"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ecssettings.html
//
type CfnComputeEnvironment_EcsSettingsProperty struct {
	// The CloudWatch Container Insights setting applied to the Amazon ECS cluster that backs this compute environment.
	//
	// After you set this property, you can't revert it to the default (unset) state in which the setting is managed outside of AWS Batch. If you remove this property after previously setting it, AWS Batch treats the omission as DISABLED, because the underlying API has no way to unset the value. Because of this, if a stack rollback would return this property to its previous unset state, AWS Batch sets it to DISABLED instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ecssettings.html#cfn-batch-computeenvironment-ecssettings-containerinsights
	//
	ContainerInsights *string `field:"optional" json:"containerInsights" yaml:"containerInsights"`
}

