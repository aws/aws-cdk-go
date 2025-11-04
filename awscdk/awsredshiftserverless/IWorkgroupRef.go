package awsredshiftserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workgroup.
// Experimental.
type IWorkgroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Workgroup resource.
	// Experimental.
	WorkgroupRef() *WorkgroupReference
}

// The jsii proxy for IWorkgroupRef
type jsiiProxy_IWorkgroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IWorkgroupRef) WorkgroupRef() *WorkgroupReference {
	var returns *WorkgroupReference
	_jsii_.Get(
		j,
		"workgroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkgroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkgroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

