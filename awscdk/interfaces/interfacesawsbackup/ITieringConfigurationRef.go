package interfacesawsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TieringConfiguration.
// Experimental.
type ITieringConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TieringConfiguration resource.
	// Experimental.
	TieringConfigurationRef() *TieringConfigurationReference
}

// The jsii proxy for ITieringConfigurationRef
type jsiiProxy_ITieringConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITieringConfigurationRef) TieringConfigurationRef() *TieringConfigurationReference {
	var returns *TieringConfigurationReference
	_jsii_.Get(
		j,
		"tieringConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITieringConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITieringConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

