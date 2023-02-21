package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IVersion interface {
	IFunction
	// Defines an alias for this version.
	// Deprecated: Calling `addAlias` on a `Version` object will cause the Alias to be replaced on every function update. Call `function.addAlias()` or `new Alias()` instead.
	AddAlias(aliasName *string, options *AliasOptions) Alias
	// The ARN of the version for Lambda@Edge.
	EdgeArn() *string
	// The underlying AWS Lambda function.
	Lambda() IFunction
	// The most recently deployed version of this function.
	Version() *string
}

// The jsii proxy for IVersion
type jsiiProxy_IVersion struct {
	jsiiProxy_IFunction
}

func (i *jsiiProxy_IVersion) AddAlias(aliasName *string, options *AliasOptions) Alias {
	if err := i.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns Alias

	_jsii_.Invoke(
		i,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVersion) EdgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

