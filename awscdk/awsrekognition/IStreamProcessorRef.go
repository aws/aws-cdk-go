package awsrekognition

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamProcessor.
// Experimental.
type IStreamProcessorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStreamProcessorRef
type jsiiProxy_IStreamProcessorRef struct {
	internal.Type__constructsIConstruct
}

