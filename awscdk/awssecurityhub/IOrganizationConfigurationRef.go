package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OrganizationConfiguration.
// Experimental.
type IOrganizationConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OrganizationConfiguration resource.
	// Experimental.
	OrganizationConfigurationRef() *OrganizationConfigurationReference
}

// The jsii proxy for IOrganizationConfigurationRef
type jsiiProxy_IOrganizationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOrganizationConfigurationRef) OrganizationConfigurationRef() *OrganizationConfigurationReference {
	var returns *OrganizationConfigurationReference
	_jsii_.Get(
		j,
		"organizationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

