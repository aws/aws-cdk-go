package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyDetector.
// Experimental.
type IAnomalyDetectorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AnomalyDetector resource.
	// Experimental.
	AnomalyDetectorRef() *AnomalyDetectorReference
}

// The jsii proxy for IAnomalyDetectorRef
type jsiiProxy_IAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAnomalyDetectorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAnomalyDetectorRef) AnomalyDetectorRef() *AnomalyDetectorReference {
	var returns *AnomalyDetectorReference
	_jsii_.Get(
		j,
		"anomalyDetectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetectorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

