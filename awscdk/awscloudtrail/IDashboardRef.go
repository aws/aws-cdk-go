package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dashboard.
// Experimental.
type IDashboardRef interface {
	constructs.IConstruct
	// A reference to a Dashboard resource.
	// Experimental.
	DashboardRef() *DashboardReference
}

// The jsii proxy for IDashboardRef
type jsiiProxy_IDashboardRef struct {
	internal.Type__constructsIConstruct
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

