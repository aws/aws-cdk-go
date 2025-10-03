package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DistributionTenant.
// Experimental.
type IDistributionTenantRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDistributionTenantRef
type jsiiProxy_IDistributionTenantRef struct {
	internal.Type__constructsIConstruct
}

