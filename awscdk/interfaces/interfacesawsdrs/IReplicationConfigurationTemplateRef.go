package interfacesawsdrs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdrs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfigurationTemplate.
// Experimental.
type IReplicationConfigurationTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReplicationConfigurationTemplate resource.
	// Experimental.
	ReplicationConfigurationTemplateRef() *ReplicationConfigurationTemplateReference
}

// The jsii proxy for IReplicationConfigurationTemplateRef
type jsiiProxy_IReplicationConfigurationTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReplicationConfigurationTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReplicationConfigurationTemplateRef) ReplicationConfigurationTemplateRef() *ReplicationConfigurationTemplateReference {
	var returns *ReplicationConfigurationTemplateReference
	_jsii_.Get(
		j,
		"replicationConfigurationTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigurationTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigurationTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

