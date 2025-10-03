package awssam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HttpApi.
// Experimental.
type IHttpApiRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHttpApiRef
type jsiiProxy_IHttpApiRef struct {
	internal.Type__constructsIConstruct
}

