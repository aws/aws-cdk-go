package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UsagePlanKey.
// Experimental.
type IUsagePlanKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUsagePlanKeyRef
type jsiiProxy_IUsagePlanKeyRef struct {
	internal.Type__constructsIConstruct
}

