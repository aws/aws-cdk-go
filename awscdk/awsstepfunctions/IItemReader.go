package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Base interface for Item Reader configurations.
type IItemReader interface {
	// Compile policy statements to provide relevent permissions to the state machine.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Render the ItemReader as JSON object.
	Render() interface{}
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	Bucket() awss3.IBucket
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems() *float64
	// The Amazon S3 API action that Step Functions must invoke depending on the specified dataset.
	Resource() *string
}

// The jsii proxy for IItemReader
type jsiiProxy_IItemReader struct {
	_ byte // padding
}

func (i *jsiiProxy_IItemReader) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		i,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IItemReader) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IItemReader) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IItemReader) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IItemReader) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

