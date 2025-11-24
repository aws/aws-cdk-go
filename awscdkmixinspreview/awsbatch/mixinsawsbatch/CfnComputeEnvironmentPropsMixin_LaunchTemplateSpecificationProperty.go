package mixinsawsbatch


// An object that represents a launch template that's associated with a compute resource.
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplateSpecificationProperty := &LaunchTemplateSpecificationProperty{
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	Overrides: []interface{}{
//   		&LaunchTemplateSpecificationOverrideProperty{
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			LaunchTemplateName: jsii.String("launchTemplateName"),
//   			TargetInstanceTypes: []*string{
//   				jsii.String("targetInstanceTypes"),
//   			},
//   			UserdataType: jsii.String("userdataType"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	UserdataType: jsii.String("userdataType"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html
//
type CfnComputeEnvironmentPropsMixin_LaunchTemplateSpecificationProperty struct {
	// The ID of the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-launchtemplateid
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// A launch template to use in place of the default launch template.
	//
	// You must specify either the launch template ID or launch template name in the request, but not both.
	//
	// You can specify up to ten (10) launch template overrides that are associated to unique instance types or families for each compute environment.
	//
	// > To unset all override templates for a compute environment, you can pass an empty array to the [UpdateComputeEnvironment.overrides](https://docs.aws.amazon.com/batch/latest/APIReference/API_UpdateComputeEnvironment.html) parameter, or not include the `overrides` parameter when submitting the `UpdateComputeEnvironment` API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// The EKS node initialization process to use.
	//
	// You only need to specify this value if you are using a custom AMI. The default value is `EKS_BOOTSTRAP_SH` . If *imageType* is a custom AMI based on EKS_AL2023 or EKS_AL2023_NVIDIA then you must choose `EKS_NODEADM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-userdatatype
	//
	UserdataType *string `field:"optional" json:"userdataType" yaml:"userdataType"`
	// The version number of the launch template, `$Default` , or `$Latest` .
	//
	// If the value is `$Default` , the default version of the launch template is used. If the value is `$Latest` , the latest version of the launch template is used.
	//
	// > If the AMI ID that's used in a compute environment is from the launch template, the AMI isn't changed when the compute environment is updated. It's only changed if the `updateToLatestImageVersion` parameter for the compute environment is set to `true` . During an infrastructure update, if either `$Default` or `$Latest` is specified, AWS Batch re-evaluates the launch template version, and it might use a different version of the launch template. This is the case even if the launch template isn't specified in the update. When updating a compute environment, changing the launch template requires an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
	//
	// Default: `$Default`
	//
	// Latest: `$Latest`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-launchtemplatespecification.html#cfn-batch-computeenvironment-launchtemplatespecification-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

