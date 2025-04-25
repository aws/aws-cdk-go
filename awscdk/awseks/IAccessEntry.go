package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
)

// Represents an access entry in an Amazon EKS cluster.
//
// An access entry defines the permissions and scope for a user or role to access an Amazon EKS cluster.
type IAccessEntry interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the access entry.
	AccessEntryArn() *string
	// The name of the access entry.
	AccessEntryName() *string
}

// The jsii proxy for IAccessEntry
type jsiiProxy_IAccessEntry struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccessEntry) AccessEntryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessEntryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) AccessEntryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessEntryName",
		&returns,
	)
	return returns
}

