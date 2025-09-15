package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationCloudWatchLoggingOption.
// Experimental.
type IApplicationCloudWatchLoggingOptionV2Ref interface {
	constructs.IConstruct
	// A reference to a ApplicationCloudWatchLoggingOption resource.
	// Experimental.
	ApplicationCloudWatchLoggingOptionRef() *ApplicationCloudWatchLoggingOptionV2Reference
}

// The jsii proxy for IApplicationCloudWatchLoggingOptionV2Ref
type jsiiProxy_IApplicationCloudWatchLoggingOptionV2Ref struct {
	internal.Type__constructsIConstruct
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

