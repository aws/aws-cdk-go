package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Detector.
// Experimental.
type IDetectorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDetectorRef
type jsiiProxy_IDetectorRef struct {
	internal.Type__constructsIConstruct
}

