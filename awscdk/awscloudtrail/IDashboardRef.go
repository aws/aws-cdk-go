package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dashboard.
// Experimental.
type IDashboardRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Dashboard resource.
	// Experimental.
	DashboardRef() *DashboardReference
}

// The jsii proxy for IDashboardRef
type jsiiProxy_IDashboardRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDashboardRef) DashboardRef() *DashboardReference {
	var returns *DashboardReference
	_jsii_.Get(
		j,
		"dashboardRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDashboardRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDashboardRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

