package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The remote access (SSH) configuration to use with your node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   nodegroupRemoteAccess := &NodegroupRemoteAccess{
//   	SshKeyName: jsii.String("sshKeyName"),
//
//   	// the properties below are optional
//   	SourceSecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-remoteaccess.html
//
// Experimental.
type NodegroupRemoteAccess struct {
	// The Amazon EC2 SSH key that provides access for SSH communication with the worker nodes in the managed node group.
	// Experimental.
	SshKeyName *string `field:"required" json:"sshKeyName" yaml:"sshKeyName"`
	// The security groups that are allowed SSH access (port 22) to the worker nodes.
	//
	// If you specify an Amazon EC2 SSH
	// key but do not specify a source security group when you create a managed node group, then port 22 on the worker
	// nodes is opened to the internet (0.0.0.0/0).
	// Default: - port 22 on the worker nodes is opened to the internet (0.0.0.0/0)
	//
	// Experimental.
	SourceSecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"sourceSecurityGroups" yaml:"sourceSecurityGroups"`
}

