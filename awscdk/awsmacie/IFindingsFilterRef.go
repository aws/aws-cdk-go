package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingsFilter.
// Experimental.
type IFindingsFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFindingsFilterRef
type jsiiProxy_IFindingsFilterRef struct {
	internal.Type__constructsIConstruct
}

