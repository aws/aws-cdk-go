package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Route 53 public hosted zone.
type IPublicHostedZone interface {
	IHostedZone
	// Grant permissions to add delegation records to this zone.
	GrantDelegation(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy for IPublicHostedZone
type jsiiProxy_IPublicHostedZone struct {
	jsiiProxy_IHostedZone
}

func (i *jsiiProxy_IPublicHostedZone) GrantDelegation(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDelegationParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDelegation",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

