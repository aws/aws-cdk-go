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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-ecssettings.html#cfn-batch-computeenvironment-ecssettings-containerinsights
	//
	ContainerInsights *string `field:"optional" json:"containerInsights" yaml:"containerInsights"`
}

