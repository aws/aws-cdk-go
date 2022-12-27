package awsbatch


// An object representing a launch template associated with a compute resource.
//
// You must specify either the launch template ID or launch template name in the request, but not both.
//
// If security groups are specified using both the `securityGroupIds` parameter of `CreateComputeEnvironment` and the launch template, the values in the `securityGroupIds` parameter of `CreateComputeEnvironment` will be used.
//
// > This object isn't applicable to jobs that are running on Fargate resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &launchTemplateSpecificationProperty{
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   	version: jsii.String("version"),
//   }
//
type CfnComputeEnvironment_LaunchTemplateSpecificationProperty struct {
	// The ID of the launch template.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version number of the launch template, `$Latest` , or `$Default` .
	//
	// If the value is `$Latest` , the latest version of the launch template is used. If the value is `$Default` , the default version of the launch template is used.
	//
	// > If the AMI ID that's used in a compute environment is from the launch template, the AMI isn't changed when the compute environment is updated. It's only changed if the `updateToLatestImageVersion` parameter for the compute environment is set to `true` . During an infrastructure update, if either `$Latest` or `$Default` is specified, AWS Batch re-evaluates the launch template version, and it might use a different version of the launch template. This is the case even if the launch template isn't specified in the update. When updating a compute environment, changing the launch template requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// Default: `$Default` .
	Version *string `field:"optional" json:"version" yaml:"version"`
}

