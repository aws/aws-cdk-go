package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudWatchAlarmTemplateGroup.
// Experimental.
type ICloudWatchAlarmTemplateGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudWatchAlarmTemplateGroup resource.
	// Experimental.
	CloudWatchAlarmTemplateGroupRef() *CloudWatchAlarmTemplateGroupReference
}

// The jsii proxy for ICloudWatchAlarmTemplateGroupRef
type jsiiProxy_ICloudWatchAlarmTemplateGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateGroupRef) CloudWatchAlarmTemplateGroupRef() *CloudWatchAlarmTemplateGroupReference {
	var returns *CloudWatchAlarmTemplateGroupReference
	_jsii_.Get(
		j,
		"cloudWatchAlarmTemplateGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

