package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Distribution.
// Experimental.
type IDistributionRef interface {
	constructs.IConstruct
	// A reference to a Distribution resource.
	// Experimental.
	DistributionRef() *DistributionReference
}

// The jsii proxy for IDistributionRef
type jsiiProxy_IDistributionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDistributionRef) DistributionRef() *DistributionReference {
	var returns *DistributionReference
	_jsii_.Get(
		j,
		"distributionRef",
		&returns,
	)
	return returns
}

