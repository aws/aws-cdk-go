package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents Realtime Log Configuration.
type IRealtimeLogConfig interface {
	interfacesawscloudfront.IRealtimeLogConfigRef
	awscdk.IResource
	// The arn of the realtime log config.
	RealtimeLogConfigArn() *string
	// The name of the realtime log config.
	RealtimeLogConfigName() *string
}

// The jsii proxy for IRealtimeLogConfig
type jsiiProxy_IRealtimeLogConfig struct {
	internal.Type__interfacesawscloudfrontIRealtimeLogConfigRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRealtimeLogConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IRealtimeLogConfig) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRealtimeLogConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRealtimeLogConfig) RealtimeLogConfigRef() *interfacesawscloudfront.RealtimeLogConfigReference {
	var returns *interfacesawscloudfront.RealtimeLogConfigReference
	_jsii_.Get(
		j,
		"realtimeLogConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRealtimeLogConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

