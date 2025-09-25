package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsagePlanKey.
// Experimental.
type IUsagePlanKeyRef interface {
	constructs.IConstruct
	// A reference to a UsagePlanKey resource.
	// Experimental.
	UsagePlanKeyRef() *UsagePlanKeyReference
}

// The jsii proxy for IUsagePlanKeyRef
type jsiiProxy_IUsagePlanKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUsagePlanKeyRef) UsagePlanKeyRef() *UsagePlanKeyReference {
	var returns *UsagePlanKeyReference
	_jsii_.Get(
		j,
		"usagePlanKeyRef",
		&returns,
	)
	return returns
}

