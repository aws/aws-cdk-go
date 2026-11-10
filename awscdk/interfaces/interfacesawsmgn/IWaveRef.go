package interfacesawsmgn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmgn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Wave.
// Experimental.
type IWaveRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Wave resource.
	// Experimental.
	WaveRef() *WaveReference
}

// The jsii proxy for IWaveRef
type jsiiProxy_IWaveRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWaveRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWaveRef) WaveRef() *WaveReference {
	var returns *WaveReference
	_jsii_.Get(
		j,
		"waveRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaveRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWaveRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

