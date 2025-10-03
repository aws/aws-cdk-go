package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConformancePack.
// Experimental.
type IOrganizationConformancePackRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOrganizationConformancePackRef
type jsiiProxy_IOrganizationConformancePackRef struct {
	internal.Type__constructsIConstruct
}

