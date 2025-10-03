package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfiguredTable.
// Experimental.
type IConfiguredTableRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfiguredTableRef
type jsiiProxy_IConfiguredTableRef struct {
	internal.Type__constructsIConstruct
}

