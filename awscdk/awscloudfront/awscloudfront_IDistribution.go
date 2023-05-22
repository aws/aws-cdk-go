package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront/internal"
)

// Interface for CloudFront distributions.
// Experimental.
type IDistribution interface {
	awscdk.IResource
	// The domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	// Experimental.
	DistributionDomainName() *string
	// The distribution ID for this distribution.
	// Experimental.
	DistributionId() *string
	// The domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	// Deprecated: - Use `distributionDomainName` instead.
	DomainName() *string
}

// The jsii proxy for IDistribution
type jsiiProxy_IDistribution struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDistribution) DistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

