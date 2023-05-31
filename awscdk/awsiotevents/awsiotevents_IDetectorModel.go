package awsiotevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotevents/internal"
)

// Represents an AWS IoT Events detector model.
// Experimental.
type IDetectorModel interface {
	awscdk.IResource
	// The name of the detector model.
	// Experimental.
	DetectorModelName() *string
}

// The jsii proxy for IDetectorModel
type jsiiProxy_IDetectorModel struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

