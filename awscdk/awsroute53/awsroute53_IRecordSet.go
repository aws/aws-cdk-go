package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53/internal"
)

// A record set.
// Experimental.
type IRecordSet interface {
	awscdk.IResource
	// The domain name of the record.
	// Experimental.
	DomainName() *string
}

// The jsii proxy for IRecordSet
type jsiiProxy_IRecordSet struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRecordSet) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

