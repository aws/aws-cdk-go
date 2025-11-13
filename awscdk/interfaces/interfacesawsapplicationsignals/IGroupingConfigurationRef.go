package interfacesawsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupingConfiguration.
// Experimental.
type IGroupingConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GroupingConfiguration resource.
	// Experimental.
	GroupingConfigurationRef() *GroupingConfigurationReference
}

// The jsii proxy for IGroupingConfigurationRef
type jsiiProxy_IGroupingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGroupingConfigurationRef) GroupingConfigurationRef() *GroupingConfigurationReference {
	var returns *GroupingConfigurationReference
	_jsii_.Get(
		j,
		"groupingConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupingConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

