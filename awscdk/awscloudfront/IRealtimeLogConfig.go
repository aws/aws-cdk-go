package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
)

// Represents Realtime Log Configuration.
type IRealtimeLogConfig interface {
	awscdk.IResource
	// The arn of the realtime log config.
	RealtimeLogConfigArn() *string
	// The name of the realtime log config.
	RealtimeLogConfigName() *string
}

// The jsii proxy for IRealtimeLogConfig
type jsiiProxy_IRealtimeLogConfig struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRealtimeLogConfig) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRealtimeLogConfig) RealtimeLogConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigName",
		&returns,
	)
	return returns
}

