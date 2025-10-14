package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregationAuthorization.
// Experimental.
type IAggregationAuthorizationRef interface {
	constructs.IConstruct
	// A reference to a AggregationAuthorization resource.
	// Experimental.
	AggregationAuthorizationRef() *AggregationAuthorizationReference
}

// The jsii proxy for IAggregationAuthorizationRef
type jsiiProxy_IAggregationAuthorizationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAggregationAuthorizationRef) AggregationAuthorizationRef() *AggregationAuthorizationReference {
	var returns *AggregationAuthorizationReference
	_jsii_.Get(
		j,
		"aggregationAuthorizationRef",
		&returns,
	)
	return returns
}

