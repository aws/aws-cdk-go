package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Collection of grant methods for a INamedHostedZoneRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namedHostedZoneRef INamedHostedZoneRef
//
//   hostedZoneGrants := awscdk.Aws_route53.HostedZoneGrants_FromHostedZone(namedHostedZoneRef)
//
type HostedZoneGrants interface {
	// Grant permissions to add delegation records to this zone.
	Delegation(grantee awsiam.IGrantable, delegationOptions *GrantDelegationOptions) awsiam.Grant
}

// The jsii proxy struct for HostedZoneGrants
type jsiiProxy_HostedZoneGrants struct {
	_ byte // padding
}

// Creates grants for INamedHostedZoneRef.
func HostedZoneGrants_FromHostedZone(hostedZone INamedHostedZoneRef) HostedZoneGrants {
	_init_.Initialize()

	if err := validateHostedZoneGrants_FromHostedZoneParameters(hostedZone); err != nil {
		panic(err)
	}
	var returns HostedZoneGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HostedZoneGrants",
		"fromHostedZone",
		[]interface{}{hostedZone},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostedZoneGrants) Delegation(grantee awsiam.IGrantable, delegationOptions *GrantDelegationOptions) awsiam.Grant {
	if err := h.validateDelegationParameters(grantee, delegationOptions); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		h,
		"delegation",
		[]interface{}{grantee, delegationOptions},
		&returns,
	)

	return returns
}

