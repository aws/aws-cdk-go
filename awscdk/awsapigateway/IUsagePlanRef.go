package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsagePlan.
// Experimental.
type IUsagePlanRef interface {
	constructs.IConstruct
	// A reference to a UsagePlan resource.
	// Experimental.
	UsagePlanRef() *UsagePlanReference
}

// The jsii proxy for IUsagePlanRef
type jsiiProxy_IUsagePlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUsagePlanRef) UsagePlanRef() *UsagePlanReference {
	var returns *UsagePlanReference
	_jsii_.Get(
		j,
		"usagePlanRef",
		&returns,
	)
	return returns
}

