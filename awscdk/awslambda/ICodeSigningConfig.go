package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
)

// A Code Signing Config.
type ICodeSigningConfig interface {
	awscdk.IResource
	// The ARN of Code Signing Config.
	CodeSigningConfigArn() *string
	// The id of Code Signing Config.
	CodeSigningConfigId() *string
}

// The jsii proxy for ICodeSigningConfig
type jsiiProxy_ICodeSigningConfig struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigId",
		&returns,
	)
	return returns
}

