package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConformancePack.
// Experimental.
type IOrganizationConformancePackRef interface {
	constructs.IConstruct
	// A reference to a OrganizationConformancePack resource.
	// Experimental.
	OrganizationConformancePackRef() *OrganizationConformancePackReference
}

// The jsii proxy for IOrganizationConformancePackRef
type jsiiProxy_IOrganizationConformancePackRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOrganizationConformancePackRef) OrganizationConformancePackRef() *OrganizationConformancePackReference {
	var returns *OrganizationConformancePackReference
	_jsii_.Get(
		j,
		"organizationConformancePackRef",
		&returns,
	)
	return returns
}

