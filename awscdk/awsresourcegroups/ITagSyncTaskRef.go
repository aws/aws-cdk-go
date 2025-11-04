package awsresourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourcegroups/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagSyncTask.
// Experimental.
type ITagSyncTaskRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TagSyncTask resource.
	// Experimental.
	TagSyncTaskRef() *TagSyncTaskReference
}

// The jsii proxy for ITagSyncTaskRef
type jsiiProxy_ITagSyncTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITagSyncTaskRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagSyncTaskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

