// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for Ec2Environment.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
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
	// The subnetSelection of the VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

