package interfacesawsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ZonalAutoshiftConfiguration.
// Experimental.
type IZonalAutoshiftConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ZonalAutoshiftConfiguration resource.
	// Experimental.
	ZonalAutoshiftConfigurationRef() *ZonalAutoshiftConfigurationReference
}

// The jsii proxy for IZonalAutoshiftConfigurationRef
type jsiiProxy_IZonalAutoshiftConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IZonalAutoshiftConfigurationRef) ZonalAutoshiftConfigurationRef() *ZonalAutoshiftConfigurationReference {
	var returns *ZonalAutoshiftConfigurationReference
	_jsii_.Get(
		j,
		"zonalAutoshiftConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IZonalAutoshiftConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IZonalAutoshiftConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

