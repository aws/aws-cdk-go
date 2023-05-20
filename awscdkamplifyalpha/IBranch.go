package awscdkamplifyalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/internal"
)

// A branch.
// Experimental.
type IBranch interface {
	awscdk.IResource
	// The name of the branch.
	// Experimental.
	BranchName() *string
}

// The jsii proxy for IBranch
type jsiiProxy_IBranch struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IBranch) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

