package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDataSync.
// Experimental.
type IResourceDataSyncRef interface {
	constructs.IConstruct
	// A reference to a ResourceDataSync resource.
	// Experimental.
	ResourceDataSyncRef() *ResourceDataSyncReference
}

// The jsii proxy for IResourceDataSyncRef
type jsiiProxy_IResourceDataSyncRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceDataSyncRef) ResourceDataSyncRef() *ResourceDataSyncReference {
	var returns *ResourceDataSyncReference
	_jsii_.Get(
		j,
		"resourceDataSyncRef",
		&returns,
	)
	return returns
}

