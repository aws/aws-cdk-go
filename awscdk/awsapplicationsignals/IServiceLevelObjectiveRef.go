package awsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLevelObjective.
// Experimental.
type IServiceLevelObjectiveRef interface {
	constructs.IConstruct
	// A reference to a ServiceLevelObjective resource.
	// Experimental.
	ServiceLevelObjectiveRef() *ServiceLevelObjectiveReference
}

// The jsii proxy for IServiceLevelObjectiveRef
type jsiiProxy_IServiceLevelObjectiveRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceLevelObjectiveRef) ServiceLevelObjectiveRef() *ServiceLevelObjectiveReference {
	var returns *ServiceLevelObjectiveReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveRef",
		&returns,
	)
	return returns
}

