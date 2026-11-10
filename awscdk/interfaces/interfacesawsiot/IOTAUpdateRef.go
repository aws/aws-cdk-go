package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OTAUpdate.
// Experimental.
type IOTAUpdateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OTAUpdate resource.
	// Experimental.
	OtaUpdateRef() *OTAUpdateReference
}

// The jsii proxy for IOTAUpdateRef
type jsiiProxy_IOTAUpdateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOTAUpdateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOTAUpdateRef) OtaUpdateRef() *OTAUpdateReference {
	var returns *OTAUpdateReference
	_jsii_.Get(
		j,
		"otaUpdateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOTAUpdateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOTAUpdateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

