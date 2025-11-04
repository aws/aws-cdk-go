package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationCloudWatchLoggingOption.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type IApplicationCloudWatchLoggingOptionV2Ref interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApplicationCloudWatchLoggingOption resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionV2Reference
}

// The jsii proxy for IApplicationCloudWatchLoggingOptionV2Ref
type jsiiProxy_IApplicationCloudWatchLoggingOptionV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionV2Ref) ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionV2Reference {
	var returns *ApplicationCloudWatchLoggingOptionV2Reference
	_jsii_.Get(
		j,
		"applicationCloudWatchLoggingOptionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionV2Ref) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationCloudWatchLoggingOptionV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

