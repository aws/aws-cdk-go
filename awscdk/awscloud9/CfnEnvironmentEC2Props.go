package awscloud9

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnvironmentEC2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentEC2Props := &CfnEnvironmentEC2Props{
//   	ImageId: jsii.String("imageId"),
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	AutomaticStopTimeMinutes: jsii.Number(123),
//   	ConnectionType: jsii.String("connectionType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OwnerArn: jsii.String("ownerArn"),
//   	Repositories: []interface{}{
//   		&RepositoryProperty{
//   			PathComponent: jsii.String("pathComponent"),
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
//
type CfnEnvironmentEC2Props struct {
	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
	//
	// To choose an AMI for the instance, you must specify a valid AMI alias or a valid AWS Systems Manager path.
	//
	// From December 04, 2023, you will be required to include the `ImageId` parameter for the `CreateEnvironmentEC2` action. This change will be reflected across all direct methods of communicating with the API, such as AWS SDK, AWS CLI and AWS CloudFormation. This change will only affect direct API consumers, and not AWS Cloud9 console users.
	//
	// Since Ubuntu 18.04 has ended standard support as of May 31, 2023, we recommend you choose Ubuntu 22.04.
	//
	// *AMI aliases*
	//
	// - Amazon Linux 2: `amazonlinux-2-x86_64`
	// - Amazon Linux 2023 (recommended): `amazonlinux-2023-x86_64`
	// - Ubuntu 18.04: `ubuntu-18.04-x86_64`
	// - Ubuntu 22.04: `ubuntu-22.04-x86_64`
	//
	// *SSM paths*
	//
	// - Amazon Linux 2: `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
	// - Amazon Linux 2023 (recommended): `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2023-x86_64`
	// - Ubuntu 18.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
	// - Ubuntu 22.04: `resolve:ssm:/aws/service/cloud9/amis/ubuntu-22.04-x86_64`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-imageid
	//
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// The type of instance to connect to the environment (for example, `t2.micro` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The number of minutes until the running instance is shut down after the environment was last used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-automaticstoptimeminutes
	//
	AutomaticStopTimeMinutes *float64 `field:"optional" json:"automaticStopTimeMinutes" yaml:"automaticStopTimeMinutes"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are `CONNECT_SSH` (default) and `CONNECT_SSM` (connected through AWS Systems Manager ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The description of the environment to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the environment owner.
	//
	// This ARN can be the ARN of any AWS Identity and Access Management principal. If this value is not specified, the ARN defaults to this environment's creator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-ownerarn
	//
	OwnerArn *string `field:"optional" json:"ownerArn" yaml:"ownerArn"`
	// Any AWS CodeCommit source code repositories to be cloned into the development environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-repositories
	//
	Repositories interface{} `field:"optional" json:"repositories" yaml:"repositories"`
	// The ID of the subnet in Amazon Virtual Private Cloud (Amazon VPC) that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html#cfn-cloud9-environmentec2-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

