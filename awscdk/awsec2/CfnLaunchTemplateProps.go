package awsec2


// Properties for defining a `CfnLaunchTemplate`.
//
// Example:
//   var cluster cluster
//
//
//   userData := `MIME-Version: 1.0
//   Content-Type: multipart/mixed; boundary="==MYBOUNDARY=="
//
//   --==MYBOUNDARY==
//   Content-Type: text/x-shellscript; charset="us-ascii"
//
//   #!/bin/bash
//   echo "Running custom user data script"
//
//   --==MYBOUNDARY==--\\
//   `
//   lt := ec2.NewCfnLaunchTemplate(this, jsii.String("LaunchTemplate"), &CfnLaunchTemplateProps{
//   	LaunchTemplateData: &LaunchTemplateDataProperty{
//   		InstanceType: jsii.String("t3.small"),
//   		UserData: awscdk.Fn_Base64(userData),
//   	},
//   })
//
//   cluster.AddNodegroupCapacity(jsii.String("extra-ng"), &NodegroupOptions{
//   	LaunchTemplateSpec: &LaunchTemplateSpec{
//   		Id: lt.ref,
//   		Version: lt.AttrLatestVersionNumber,
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html
//
type CfnLaunchTemplateProps struct {
	// The information for the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatedata
	//
	LaunchTemplateData interface{} `field:"required" json:"launchTemplateData" yaml:"launchTemplateData"`
	// A name for the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The tags to apply to the launch template on creation.
	//
	// To tag the launch template, the resource type must be `launch-template` .
	//
	// To specify the tags for the resources that are created when an instance is launched, you must use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// A description for the first version of the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

