package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalyticsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationCloudWatchLoggingOption.
// Experimental.
type IApplicationCloudWatchLoggingOptionRef interface {
	constructs.IConstruct
	// A reference to a ApplicationCloudWatchLoggingOption resource.
	// Experimental.
	ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionReference
}

// The jsii proxy for IApplicationCloudWatchLoggingOptionRef
type jsiiProxy_IApplicationCloudWatchLoggingOptionRef struct {
	internal.Type__constructsIConstruct
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

