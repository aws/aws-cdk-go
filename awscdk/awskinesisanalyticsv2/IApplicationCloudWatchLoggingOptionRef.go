package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalyticsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationCloudWatchLoggingOption.
// Experimental.
type IApplicationCloudWatchLoggingOptionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApplicationCloudWatchLoggingOption resource.
	// Experimental.
	ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionReference
}

// The jsii proxy for IApplicationCloudWatchLoggingOptionRef
type jsiiProxy_IApplicationCloudWatchLoggingOptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

