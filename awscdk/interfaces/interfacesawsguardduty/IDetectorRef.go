package interfacesawsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Detector.
// Experimental.
type IDetectorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Detector resource.
	// Experimental.
	DetectorRef() *DetectorReference
}

// The jsii proxy for IDetectorRef
type jsiiProxy_IDetectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDetectorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDetectorRef) DetectorRef() *DetectorReference {
	var returns *DetectorReference
	_jsii_.Get(
		j,
		"detectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDetectorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDetectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

