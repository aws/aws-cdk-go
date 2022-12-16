// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha


// The image ID used for creating an Amazon EC2 environment.
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
type ImageId string

const (
	// Create using Amazon Linux 2.
	// Experimental.
	ImageId_AMAZON_LINUX_2 ImageId = "AMAZON_LINUX_2"
	// Create using Ubunut 18.04.
	// Experimental.
	ImageId_UBUNTU_18_04 ImageId = "UBUNTU_18_04"
)

