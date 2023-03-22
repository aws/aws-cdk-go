// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcloud9alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
)

// The class for different repository providers.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &RepositoryProps{
//   	RepositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.Repository_FromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &Ec2EnvironmentProps{
//   	Vpc: Vpc,
//   	ClonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository_FromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository_*FromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
//   })
//
// Experimental.
type CloneRepository interface {
	// Experimental.
	PathComponent() *string
	// Experimental.
	RepositoryUrl() *string
}

// The jsii proxy struct for CloneRepository
type jsiiProxy_CloneRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CloneRepository) PathComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloneRepository) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}


// import repository to cloud9 environment from AWS CodeCommit.
// Experimental.
func CloneRepository_FromCodeCommit(repository awscodecommit.IRepository, path *string) CloneRepository {
	_init_.Initialize()

	if err := validateCloneRepository_FromCodeCommitParameters(repository, path); err != nil {
		panic(err)
	}
	var returns CloneRepository

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.CloneRepository",
		"fromCodeCommit",
		[]interface{}{repository, path},
		&returns,
	)

	return returns
}

