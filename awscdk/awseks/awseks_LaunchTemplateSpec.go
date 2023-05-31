package awseks


// Launch template property specification.
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
// Experimental.
type LaunchTemplateSpec struct {
	// The Launch template ID.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

