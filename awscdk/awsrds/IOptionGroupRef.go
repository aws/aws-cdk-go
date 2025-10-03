package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptionGroup.
// Experimental.
type IOptionGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOptionGroupRef
type jsiiProxy_IOptionGroupRef struct {
	internal.Type__constructsIConstruct
}

