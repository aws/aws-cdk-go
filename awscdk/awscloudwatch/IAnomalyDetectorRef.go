package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyDetector.
// Experimental.
type IAnomalyDetectorRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AnomalyDetector resource.
	// Experimental.
	AnomalyDetectorRef() *AnomalyDetectorReference
}

// The jsii proxy for IAnomalyDetectorRef
type jsiiProxy_IAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAnomalyDetectorRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

