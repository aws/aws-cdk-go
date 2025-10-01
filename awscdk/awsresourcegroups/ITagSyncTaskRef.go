package awsresourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourcegroups/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagSyncTask.
// Experimental.
type ITagSyncTaskRef interface {
	constructs.IConstruct
	// A reference to a TagSyncTask resource.
	// Experimental.
	TagSyncTaskRef() *TagSyncTaskReference
}

// The jsii proxy for ITagSyncTaskRef
type jsiiProxy_ITagSyncTaskRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITagSyncTaskRef) TagSyncTaskRef() *TagSyncTaskReference {
	var returns *TagSyncTaskReference
	_jsii_.Get(
		j,
		"tagSyncTaskRef",
		&returns,
	)
	return returns
}

