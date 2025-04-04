package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A prefix list.
type IPrefixList interface {
	awscdk.IResource
	// The ID of the prefix list.
	PrefixListId() *string
}

// The jsii proxy for IPrefixList
type jsiiProxy_IPrefixList struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPrefixList) PrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixListId",
		&returns,
	)
	return returns
}

