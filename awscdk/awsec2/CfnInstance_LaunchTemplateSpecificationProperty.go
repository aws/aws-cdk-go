package awsec2


// Specifies a launch template to use when launching an Amazon EC2 instance.
//
// You must specify the following:
//
// - The ID or the name of the launch template, but not both.
// - The version of the launch template.
//
// For information about creating a launch template, see [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) and [Create a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) in the *Amazon EC2 User Guide* . For example launch templates, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#aws-resource-ec2-launchtemplate--examples) for `AWS::EC2::LaunchTemplate` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &LaunchTemplateSpecificationProperty{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-launchtemplatespecification.html
//
type CfnInstance_LaunchTemplateSpecificationProperty struct {
	// The version number of the launch template. You must specify this property.
	//
	// To specify the default version of the template, use the `Fn::GetAtt` intrinsic function to retrieve the `DefaultVersionNumber` attribute of the launch template. To specify the latest version of the template, use `Fn::GetAtt` to retrieve the `LatestVersionNumber` attribute. For more information, see [AWS::EC2:LaunchTemplate return values for Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#aws-resource-ec2-launchtemplate-return-values-fn--getatt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-launchtemplatespecification.html#cfn-ec2-instance-launchtemplatespecification-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	//
	// You must specify either the launch template ID or the launch template name, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-launchtemplatespecification.html#cfn-ec2-instance-launchtemplatespecification-launchtemplateid
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	//
	// You must specify either the launch template ID or the launch template name, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-launchtemplatespecification.html#cfn-ec2-instance-launchtemplatespecification-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

