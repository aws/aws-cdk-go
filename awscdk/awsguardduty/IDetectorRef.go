package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Detector.
// Experimental.
type IDetectorRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Detector resource.
	// Experimental.
	DetectorRef() *DetectorReference
}

// The jsii proxy for IDetectorRef
type jsiiProxy_IDetectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IDetectorRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

