package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for CloudFront distributions.
type IDistribution interface {
	interfacesawscloudfront.IDistributionRef
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
	internal.Type__interfacesawscloudfrontIDistributionRef
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

func (i *jsiiProxy_IDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
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

func (j *jsiiProxy_IDistribution) DistributionRef() *interfacesawscloudfront.DistributionReference {
	var returns *interfacesawscloudfront.DistributionReference
	_jsii_.Get(
		j,
		"distributionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

