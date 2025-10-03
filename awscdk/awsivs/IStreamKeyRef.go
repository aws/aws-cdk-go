package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamKey.
// Experimental.
type IStreamKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStreamKeyRef
type jsiiProxy_IStreamKeyRef struct {
	internal.Type__constructsIConstruct
}

