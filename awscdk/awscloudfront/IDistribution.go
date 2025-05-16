package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Interface for CloudFront distributions.
type IDistribution interface {
	awscdk.IResource
	// Adds an IAM policy statement associated with this distribution to an IAM principal's policy.
	Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant to create invalidations for this bucket to an IAM principal (Role/Group/User).
	GrantCreateInvalidation(identity awsiam.IGrantable) awsiam.Grant
	// The distribution ARN for this distribution.
	DistributionArn() *string
	// The domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DistributionDomainName() *string
	// The distribution ID for this distribution.
	DistributionId() *string
}

// The jsii proxy for IDistribution
type jsiiProxy_IDistribution struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDistribution) Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDistribution) GrantCreateInvalidation(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantCreateInvalidationParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantCreateInvalidation",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDistribution) DistributionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionArn",
		&returns,
	)
	return returns
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

