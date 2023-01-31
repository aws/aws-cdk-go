package awscodepipelineactions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions/internal"
)

// A Jenkins provider.
//
// If you want to create a new Jenkins provider managed alongside your CDK code,
// instantiate the {@link JenkinsProvider} class directly.
//
// If you want to reference an already registered provider,
// use the {@link JenkinsProvider#fromJenkinsProviderAttributes} method.
// Experimental.
type IJenkinsProvider interface {
	awscdk.IConstruct
	// Experimental.
	ProviderName() *string
	// Experimental.
	ServerUrl() *string
	// Experimental.
	Version() *string
}

// The jsii proxy for IJenkinsProvider
type jsiiProxy_IJenkinsProvider struct {
	internal.Type__awscdkIConstruct
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

