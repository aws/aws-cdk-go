package awsresiliencehub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsresiliencehub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResiliencyPolicy.
// Experimental.
type IResiliencyPolicyRef interface {
	constructs.IConstruct
	// A reference to a ResiliencyPolicy resource.
	// Experimental.
	ResiliencyPolicyRef() *ResiliencyPolicyReference
}

// The jsii proxy for IResiliencyPolicyRef
type jsiiProxy_IResiliencyPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResiliencyPolicyRef) ResiliencyPolicyRef() *ResiliencyPolicyReference {
	var returns *ResiliencyPolicyReference
	_jsii_.Get(
		j,
		"resiliencyPolicyRef",
		&returns,
	)
	return returns
}

