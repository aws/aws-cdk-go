package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents a provider for an AWS CloudFormation custom resources.
// Deprecated: use `core.ICustomResourceProvider`
type ICustomResourceProvider interface {
	// Called when this provider is used by a `CustomResource`.
	//
	// Returns: provider configuration.
	// Deprecated: use `core.ICustomResourceProvider`
	Bind(scope awscdk.Construct) *CustomResourceProviderConfig
}

// The jsii proxy for ICustomResourceProvider
type jsiiProxy_ICustomResourceProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_ICustomResourceProvider) Bind(scope awscdk.Construct) *CustomResourceProviderConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CustomResourceProviderConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

