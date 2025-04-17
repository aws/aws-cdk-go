package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
)

// Represents a user pool group.
type IUserPoolGroup interface {
	awscdk.IResource
	// The user group name.
	GroupName() *string
}

// The jsii proxy for IUserPoolGroup
type jsiiProxy_IUserPoolGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUserPoolGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

