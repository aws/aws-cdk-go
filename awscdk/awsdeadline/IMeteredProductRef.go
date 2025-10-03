package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MeteredProduct.
// Experimental.
type IMeteredProductRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMeteredProductRef
type jsiiProxy_IMeteredProductRef struct {
	internal.Type__constructsIConstruct
}

