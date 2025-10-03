package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupVersion.
// Experimental.
type IGroupVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGroupVersionRef
type jsiiProxy_IGroupVersionRef struct {
	internal.Type__constructsIConstruct
}

