package awscloud9

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEnvironmentEC2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentEC2Props := &cfnEnvironmentEC2Props{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	automaticStopTimeMinutes: jsii.Number(123),
//   	connectionType: jsii.String("connectionType"),
//   	description: jsii.String("description"),
//   	imageId: jsii.String("imageId"),
//   	name: jsii.String("name"),
//   	ownerArn: jsii.String("ownerArn"),
//   	repositories: []interface{}{
//   		&repositoryProperty{
//   			pathComponent: jsii.String("pathComponent"),
//   			repositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	subnetId: jsii.String("subnetId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEnvironmentEC2Props struct {
	// The type of instance to connect to the environment (for example, `t2.micro` ).
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The number of minutes until the running instance is shut down after the environment was last used.
	AutomaticStopTimeMinutes *float64 `field:"optional" json:"automaticStopTimeMinutes" yaml:"automaticStopTimeMinutes"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are `CONNECT_SSH` (default) and `CONNECT_SSM` (connected through AWS Systems Manager ).
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The description of the environment to create.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
	//
	// To choose an AMI for the instance, you must specify a valid AMI alias or a valid AWS Systems Manager path.
	//
	// The default AMI is used if the parameter isn't explicitly assigned a value in the request.
	//
	// *AMI aliases*
	//
	// - *Amazon Linux (default): `amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `ubuntu-18.04-x86_64`
	//
	// *SSM paths*
	//
	// - *Amazon Linux (default): `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`*
	// - Amazon Linux 2: `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
	// - Ubuntu 18.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The name of the environment.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the environment owner.
	//
	// This ARN can be the ARN of any AWS Identity and Access Management principal. If this value is not specified, the ARN defaults to this environment's creator.
	OwnerArn *string `field:"optional" json:"ownerArn" yaml:"ownerArn"`
	// Any AWS CodeCommit source code repositories to be cloned into the development environment.
	Repositories interface{} `field:"optional" json:"repositories" yaml:"repositories"`
	// The ID of the subnet in Amazon Virtual Private Cloud (Amazon VPC) that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

