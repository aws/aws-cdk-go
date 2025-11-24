package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// A container repository used to distribute container images in EC2 Image Builder.
//
// Example:
//   sourceRepo := ecr.Repository_FromRepositoryName(this, jsii.String("SourceRepo"), jsii.String("my-base-image"))
//   targetRepo := ecr.Repository_FromRepositoryName(this, jsii.String("TargetRepo"), jsii.String("my-container-repo"))
//
//   containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("EcrContainerRecipe"), &ContainerRecipeProps{
//   	BaseImage: imagebuilder.BaseContainerImage_FromEcr(sourceRepo, jsii.String("1.0.0")),
//   	TargetRepository: imagebuilder.Repository_FromEcr(targetRepo),
//   })
//
// Experimental.
type Repository interface {
	// The name of the container repository where the output container image is stored.
	// Experimental.
	RepositoryName() *string
	// The service in which the container repository is hosted.
	// Experimental.
	Service() RepositoryService
}

// The jsii proxy struct for Repository
type jsiiProxy_Repository struct {
	_ byte // padding
}

func (j *jsiiProxy_Repository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Service() RepositoryService {
	var returns RepositoryService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}


// Experimental.
func NewRepository_Override(r Repository) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.Repository",
		nil, // no parameters
		r,
	)
}

// The ECR repository to use as the target container repository.
// Experimental.
func Repository_FromEcr(repository awsecr.IRepository) Repository {
	_init_.Initialize()

	if err := validateRepository_FromEcrParameters(repository); err != nil {
		panic(err)
	}
	var returns Repository

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.Repository",
		"fromEcr",
		[]interface{}{repository},
		&returns,
	)

	return returns
}

