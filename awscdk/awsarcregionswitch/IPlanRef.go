package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsarcregionswitch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Plan.
// Experimental.
type IPlanRef interface {
	constructs.IConstruct
	// A reference to a Plan resource.
	// Experimental.
	PlanRef() *PlanReference
}

// The jsii proxy for IPlanRef
type jsiiProxy_IPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPlanRef) PlanRef() *PlanReference {
	var returns *PlanReference
	_jsii_.Get(
		j,
		"planRef",
		&returns,
	)
	return returns
}

