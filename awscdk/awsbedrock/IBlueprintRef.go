package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Blueprint.
// Experimental.
type IBlueprintRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBlueprintRef
type jsiiProxy_IBlueprintRef struct {
	internal.Type__constructsIConstruct
}

