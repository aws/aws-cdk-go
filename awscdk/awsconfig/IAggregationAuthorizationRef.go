package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregationAuthorization.
// Experimental.
type IAggregationAuthorizationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAggregationAuthorizationRef
type jsiiProxy_IAggregationAuthorizationRef struct {
	internal.Type__constructsIConstruct
}

