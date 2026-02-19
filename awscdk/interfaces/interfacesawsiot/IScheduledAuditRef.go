package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAudit.
// Experimental.
type IScheduledAuditRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ScheduledAudit resource.
	// Experimental.
	ScheduledAuditRef() *ScheduledAuditReference
}

// The jsii proxy for IScheduledAuditRef
type jsiiProxy_IScheduledAuditRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IScheduledAuditRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IScheduledAuditRef) ScheduledAuditRef() *ScheduledAuditReference {
	var returns *ScheduledAuditReference
	_jsii_.Get(
		j,
		"scheduledAuditRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledAuditRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduledAuditRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

