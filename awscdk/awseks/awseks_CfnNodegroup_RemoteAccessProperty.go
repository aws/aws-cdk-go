package awseks


// An object representing the remote access configuration for the managed node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteAccessProperty := &remoteAccessProperty{
//   	ec2SshKey: jsii.String("ec2SshKey"),
//
//   	// the properties below are optional
//   	sourceSecurityGroups: []*string{
//   		jsii.String("sourceSecurityGroups"),
//   	},
//   }
//
type CfnNodegroup_RemoteAccessProperty struct {
	// The Amazon EC2 SSH key name that provides access for SSH communication with the nodes in the managed node group.
	//
	// For more information, see [Amazon EC2 key pairs and Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Linux Instances* . For Windows, an Amazon EC2 SSH key is used to obtain the RDP password. For more information, see [Amazon EC2 key pairs and Windows instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-key-pairs.html) in the *Amazon Elastic Compute Cloud User Guide for Windows Instances* .
	Ec2SshKey *string `field:"required" json:"ec2SshKey" yaml:"ec2SshKey"`
	// The security group IDs that are allowed SSH access (port 22) to the nodes.
	//
	// For Windows, the port is 3389. If you specify an Amazon EC2 SSH key but don't specify a source security group when you create a managed node group, then the port on the nodes is opened to the internet ( `0.0.0.0/0` ). For more information, see [Security Groups for Your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the *Amazon Virtual Private Cloud User Guide* .
	SourceSecurityGroups *[]*string `field:"optional" json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

