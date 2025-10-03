package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConformancePack.
// Experimental.
type IConformancePackRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConformancePackRef
type jsiiProxy_IConformancePackRef struct {
	internal.Type__constructsIConstruct
}

