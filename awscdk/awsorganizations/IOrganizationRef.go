package awsorganizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Organization.
// Experimental.
type IOrganizationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Organization resource.
	// Experimental.
	OrganizationRef() *OrganizationReference
}

// The jsii proxy for IOrganizationRef
type jsiiProxy_IOrganizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOrganizationRef) OrganizationRef() *OrganizationReference {
	var returns *OrganizationReference
	_jsii_.Get(
		j,
		"organizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

