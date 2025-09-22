package awsredshiftserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workgroup.
// Experimental.
type IWorkgroupRef interface {
	constructs.IConstruct
	// A reference to a Workgroup resource.
	// Experimental.
	WorkgroupRef() *WorkgroupReference
}

// The jsii proxy for IWorkgroupRef
type jsiiProxy_IWorkgroupRef struct {
	internal.Type__constructsIConstruct
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

