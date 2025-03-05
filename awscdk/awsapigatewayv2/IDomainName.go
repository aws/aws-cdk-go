package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
)

// Represents an APIGatewayV2 DomainName.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html
//
type IDomainName interface {
	awscdk.IResource
	// The custom domain name.
	Name() *string
	// The domain name associated with the regional endpoint for this custom domain name.
	RegionalDomainName() *string
	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	RegionalHostedZoneId() *string
}

// The jsii proxy for IDomainName
type jsiiProxy_IDomainName struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDomainName) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainName) RegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalHostedZoneId",
		&returns,
	)
	return returns
}

