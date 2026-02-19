package interfacesawskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisanalyticsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationCloudWatchLoggingOption.
// Experimental.
type IApplicationCloudWatchLoggingOptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationCloudWatchLoggingOption resource.
	// Experimental.
	ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionReference
}

// The jsii proxy for IApplicationCloudWatchLoggingOptionRef
type jsiiProxy_IApplicationCloudWatchLoggingOptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApplicationCloudWatchLoggingOptionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionRef) ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionReference {
	var returns *ApplicationCloudWatchLoggingOptionReference
	_jsii_.Get(
		j,
		"applicationCloudWatchLoggingOptionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

