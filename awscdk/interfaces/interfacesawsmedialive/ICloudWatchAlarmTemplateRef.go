package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudWatchAlarmTemplate.
// Experimental.
type ICloudWatchAlarmTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudWatchAlarmTemplate resource.
	// Experimental.
	CloudWatchAlarmTemplateRef() *CloudWatchAlarmTemplateReference
}

// The jsii proxy for ICloudWatchAlarmTemplateRef
type jsiiProxy_ICloudWatchAlarmTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateRef) CloudWatchAlarmTemplateRef() *CloudWatchAlarmTemplateReference {
	var returns *CloudWatchAlarmTemplateReference
	_jsii_.Get(
		j,
		"cloudWatchAlarmTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

