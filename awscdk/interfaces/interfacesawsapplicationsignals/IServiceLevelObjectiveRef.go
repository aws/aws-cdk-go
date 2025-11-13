package interfacesawsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLevelObjective.
// Experimental.
type IServiceLevelObjectiveRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceLevelObjective resource.
	// Experimental.
	ServiceLevelObjectiveRef() *ServiceLevelObjectiveReference
}

// The jsii proxy for IServiceLevelObjectiveRef
type jsiiProxy_IServiceLevelObjectiveRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IServiceLevelObjectiveRef) ServiceLevelObjectiveRef() *ServiceLevelObjectiveReference {
	var returns *ServiceLevelObjectiveReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceLevelObjectiveRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceLevelObjectiveRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

