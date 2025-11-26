package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3-specific mixin for enabling versioning.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//
//
//   bucket := s3.NewCfnBucket(scope, jsii.String("MyBucket")).with(awscdkmixinspreview.NewEnableVersioning()).with(awscdkmixinspreview.NewAutoDeleteObjects())
//
// Experimental.
type EnableVersioning interface {
	core.IMixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for EnableVersioning
type jsiiProxy_EnableVersioning struct {
	internal.Type__coreIMixin
}

// Experimental.
func NewEnableVersioning() EnableVersioning {
	_init_.Initialize()

	j := jsiiProxy_EnableVersioning{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.EnableVersioning",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEnableVersioning_Override(e EnableVersioning) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.EnableVersioning",
		nil, // no parameters
		e,
	)
}

func (e *jsiiProxy_EnableVersioning) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := e.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnableVersioning) Supports(construct constructs.IConstruct) *bool {
	if err := e.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		e,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

