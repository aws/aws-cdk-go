package awscodepipelineactions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Jenkins provider.
//
// If you want to create a new Jenkins provider managed alongside your CDK code,
// instantiate the `JenkinsProvider` class directly.
//
// If you want to reference an already registered provider,
// use the `JenkinsProvider#fromJenkinsProviderAttributes` method.
type IJenkinsProvider interface {
	constructs.IConstruct
	ProviderName() *string
	ServerUrl() *string
	Version() *string
}

// The jsii proxy for IJenkinsProvider
type jsiiProxy_IJenkinsProvider struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJenkinsProvider) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJenkinsProvider) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJenkinsProvider) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

