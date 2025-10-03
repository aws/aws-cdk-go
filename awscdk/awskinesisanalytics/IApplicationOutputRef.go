package awskinesisanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Experimental.
type IApplicationOutputRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationOutputRef
type jsiiProxy_IApplicationOutputRef struct {
	internal.Type__constructsIConstruct
}

