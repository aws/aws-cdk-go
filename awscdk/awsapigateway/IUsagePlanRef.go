package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsagePlan.
// Experimental.
type IUsagePlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUsagePlanRef
type jsiiProxy_IUsagePlanRef struct {
	internal.Type__constructsIConstruct
}

