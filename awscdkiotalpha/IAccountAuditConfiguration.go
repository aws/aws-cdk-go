package awscdkiotalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
)

// Represents AWS IoT Audit Configuration.
// Experimental.
type IAccountAuditConfiguration interface {
	awscdk.IResource
	// The account ID.
	// Experimental.
	AccountId() *string
}

// The jsii proxy for IAccountAuditConfiguration
type jsiiProxy_IAccountAuditConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccountAuditConfiguration) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

