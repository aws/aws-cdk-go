package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsageProfile.
// Experimental.
type IUsageProfileRef interface {
	constructs.IConstruct
	// A reference to a UsageProfile resource.
	// Experimental.
	UsageProfileRef() *UsageProfileReference
}

// The jsii proxy for IUsageProfileRef
type jsiiProxy_IUsageProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUsageProfileRef) UsageProfileRef() *UsageProfileReference {
	var returns *UsageProfileReference
	_jsii_.Get(
		j,
		"usageProfileRef",
		&returns,
	)
	return returns
}

