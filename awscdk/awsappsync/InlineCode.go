package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// AppSync function code from an inline string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineCode := awscdk.Aws_appsync.NewInlineCode(jsii.String("code"))
//
type InlineCode interface {
	Code
	// Bind source code to an AppSync Function or resolver.
	Bind(_scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	if err := validateNewInlineCodeParameters(code); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Loads the function code from a local disk path.
func InlineCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateInlineCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.InlineCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Inline code for AppSync function.
//
// Returns: `InlineCode` with inline code.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateInlineCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineCode) Bind(_scope constructs.Construct) *CodeConfig {
	if err := i.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

