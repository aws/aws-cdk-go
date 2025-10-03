package awssecuritylake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AwsLogSource.
// Experimental.
type IAwsLogSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAwsLogSourceRef
type jsiiProxy_IAwsLogSourceRef struct {
	internal.Type__constructsIConstruct
}

