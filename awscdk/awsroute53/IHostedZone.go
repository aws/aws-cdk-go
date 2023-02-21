package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
)

// Imported or created hosted zone.
type IHostedZone interface {
	awscdk.IResource
	// ARN of this hosted zone, such as arn:${Partition}:route53:::hostedzone/${Id}.
	HostedZoneArn() *string
	// ID of this hosted zone, such as "Z23ABC4XYZL05B".
	HostedZoneId() *string
	// Returns the set of name servers for the specific hosted zone. For example: ns1.example.com.
	//
	// This attribute will be undefined for private hosted zones or hosted zones imported from another stack.
	HostedZoneNameServers() *[]*string
	// FQDN of this hosted zone.
	ZoneName() *string
}

// The jsii proxy for IHostedZone
type jsiiProxy_IHostedZone struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IHostedZone) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) HostedZoneNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostedZoneNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZone) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}

