// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for Ec2Environment.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var vpc vpc
//
//
//   user := iam.NewUser(this, jsii.String("user"))
//   user.addManagedPolicy(iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("AWSCloud9Administrator")))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//
//   	owner: cloud9.owner.user(user),
//   })
//
// Experimental.
type Ec2EnvironmentProps struct {
	// The image ID used for creating an Amazon EC2 environment.
	// Experimental.
	ImageId ImageId `field:"required" json:"imageId" yaml:"imageId"`
	// The VPC that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The AWS CodeCommit repository to be cloned.
	// Experimental.
	ClonedRepositories *[]CloneRepository `field:"optional" json:"clonedRepositories" yaml:"clonedRepositories"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are: CONNECT_SSH (default) and CONNECT_SSM (connected through AWS Systems Manager).
	// Experimental.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Description of the environment.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the environment.
	// Experimental.
	Ec2EnvironmentName *string `field:"optional" json:"ec2EnvironmentName" yaml:"ec2EnvironmentName"`
	// The type of instance to connect to the environment.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Owner of the environment.
	//
	// The owner has full control of the environment and can invite additional members.
	// Experimental.
	Owner Owner `field:"optional" json:"owner" yaml:"owner"`
	// The subnetSelection of the VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

