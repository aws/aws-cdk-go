package awsiotevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DetectorModel.
// Experimental.
type IDetectorModelRef interface {
	constructs.IConstruct
	// A reference to a DetectorModel resource.
	// Experimental.
	DetectorModelRef() *DetectorModelReference
}

// The jsii proxy for IDetectorModelRef
type jsiiProxy_IDetectorModelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDetectorModelRef) DetectorModelRef() *DetectorModelReference {
	var returns *DetectorModelReference
	_jsii_.Get(
		j,
		"detectorModelRef",
		&returns,
	)
	return returns
}

