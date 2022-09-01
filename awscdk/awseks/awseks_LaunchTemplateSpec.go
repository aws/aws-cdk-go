package awseks


// Launch template property specification.
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
type LaunchTemplateSpec struct {
	// The Launch template ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	Version *string `field:"optional" json:"version" yaml:"version"`
}

