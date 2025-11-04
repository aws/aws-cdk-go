package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DistributionTenant.
// Experimental.
type IDistributionTenantRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DistributionTenant resource.
	// Experimental.
	DistributionTenantRef() *DistributionTenantReference
}

// The jsii proxy for IDistributionTenantRef
type jsiiProxy_IDistributionTenantRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDistributionTenantRef) DistributionTenantRef() *DistributionTenantReference {
	var returns *DistributionTenantReference
	_jsii_.Get(
		j,
		"distributionTenantRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistributionTenantRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistributionTenantRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

