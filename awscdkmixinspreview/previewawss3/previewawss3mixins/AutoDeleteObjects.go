package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific mixin for auto-deleting objects.
//
// Example:
//   bucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
//   awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewAutoDeleteObjects())
//
// Experimental.
type AutoDeleteObjects interface {
	core.IMixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for AutoDeleteObjects
type jsiiProxy_AutoDeleteObjects struct {
	internal.Type__coreIMixin
}

// Experimental.
func NewAutoDeleteObjects() AutoDeleteObjects {
	_init_.Initialize()

	j := jsiiProxy_AutoDeleteObjects{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.AutoDeleteObjects",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAutoDeleteObjects_Override(a AutoDeleteObjects) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.AutoDeleteObjects",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AutoDeleteObjects) ApplyTo(construct constructs.IConstruct) {
	if err := a.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyTo",
		[]interface{}{construct},
	)
}

func (a *jsiiProxy_AutoDeleteObjects) Supports(construct constructs.IConstruct) *bool {
	if err := a.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

