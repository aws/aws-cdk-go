package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BillingGroup.
// Experimental.
type IBillingGroupRef interface {
	constructs.IConstruct
	// A reference to a BillingGroup resource.
	// Experimental.
	BillingGroupRef() *BillingGroupReference
}

// The jsii proxy for IBillingGroupRef
type jsiiProxy_IBillingGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBillingGroupRef) BillingGroupRef() *BillingGroupReference {
	var returns *BillingGroupReference
	_jsii_.Get(
		j,
		"billingGroupRef",
		&returns,
	)
	return returns
}

