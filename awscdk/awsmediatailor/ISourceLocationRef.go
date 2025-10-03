package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceLocation.
// Experimental.
type ISourceLocationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISourceLocationRef
type jsiiProxy_ISourceLocationRef struct {
	internal.Type__constructsIConstruct
}

