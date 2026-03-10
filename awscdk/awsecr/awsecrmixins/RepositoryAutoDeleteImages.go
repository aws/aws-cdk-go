package awsecrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/awsecrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// ECR-specific Mixin to force-delete all images from a repository when the repository is removed from the stack or when the stack is deleted.
//
// Sets the `emptyOnDelete` property on the repository.
//
// Example:
//   ecr.NewCfnRepository(this, jsii.String("Repo")).With(mixins.NewRepositoryAutoDeleteImages())
//
type RepositoryAutoDeleteImages interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for RepositoryAutoDeleteImages
type jsiiProxy_RepositoryAutoDeleteImages struct {
	internal.Type__awscdkMixin
}

func NewRepositoryAutoDeleteImages() RepositoryAutoDeleteImages {
	_init_.Initialize()

	j := jsiiProxy_RepositoryAutoDeleteImages{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.mixins.RepositoryAutoDeleteImages",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewRepositoryAutoDeleteImages_Override(r RepositoryAutoDeleteImages) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.mixins.RepositoryAutoDeleteImages",
		nil, // no parameters
		r,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func RepositoryAutoDeleteImages_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryAutoDeleteImages_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.mixins.RepositoryAutoDeleteImages",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryAutoDeleteImages) ApplyTo(construct constructs.IConstruct) {
	if err := r.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyTo",
		[]interface{}{construct},
	)
}

func (r *jsiiProxy_RepositoryAutoDeleteImages) Supports(construct constructs.IConstruct) *bool {
	if err := r.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

