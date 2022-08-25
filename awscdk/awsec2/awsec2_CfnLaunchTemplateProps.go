package awsec2


// Properties for defining a `CfnLaunchTemplate`.
//
// Example:
//   var cluster cluster
//
//
//   userData := "MIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"==MYBOUNDARY==\"\n\n--==MYBOUNDARY==\nContent-Type: text/x-shellscript; charset=\"us-ascii\"\n\n#!/bin/bash\necho \"Running custom user data script\"\n\n--==MYBOUNDARY==--\\\n"
//   lt := ec2.NewCfnLaunchTemplate(this, jsii.String("LaunchTemplate"), &cfnLaunchTemplateProps{
//   	launchTemplateData: &launchTemplateDataProperty{
//   		instanceType: jsii.String("t3.small"),
//   		userData: awscdk.Fn.base64(userData),
//   	},
//   })
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng"), &nodegroupOptions{
//   	launchTemplateSpec: &launchTemplateSpec{
//   		id: lt.ref,
//   		version: lt.attrLatestVersionNumber,
//   	},
//   })
//
type CfnLaunchTemplateProps struct {
	// The information for the launch template.
	LaunchTemplateData interface{} `field:"required" json:"launchTemplateData" yaml:"launchTemplateData"`
	// A name for the launch template.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The tags to apply to the launch template during creation.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// `AWS::EC2::LaunchTemplate.VersionDescription`.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

