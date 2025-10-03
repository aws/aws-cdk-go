package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RunGroup.
// Experimental.
type IRunGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRunGroupRef
type jsiiProxy_IRunGroupRef struct {
	internal.Type__constructsIConstruct
}

