package awssagemaker


// Configuration parameters specifying IAM roles assumed by SageMaker's execution role and cluster instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   emrSettingsProperty := &EmrSettingsProperty{
//   	AssumableRoleArns: []*string{
//   		jsii.String("assumableRoleArns"),
//   	},
//   	ExecutionRoleArns: []*string{
//   		jsii.String("executionRoleArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-emrsettings.html
//
type CfnUserProfilePropsMixin_EmrSettingsProperty struct {
	// An array of Amazon Resource Names (ARNs) of the IAM roles that the execution role of SageMaker can assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-emrsettings.html#cfn-sagemaker-userprofile-emrsettings-assumablerolearns
	//
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// An array of ARNs of IAM roles used by EMR cluster instances or job execution environments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-emrsettings.html#cfn-sagemaker-userprofile-emrsettings-executionrolearns
	//
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}

