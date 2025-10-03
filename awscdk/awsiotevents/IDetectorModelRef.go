package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DetectorModel.
// Experimental.
type IDetectorModelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDetectorModelRef
type jsiiProxy_IDetectorModelRef struct {
	internal.Type__constructsIConstruct
}

