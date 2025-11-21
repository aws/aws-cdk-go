package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Distribution Configuration.
// Experimental.
type IDistributionConfiguration interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the distribution configuration.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the distribution configuration.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the distribution configuration.
	// Experimental.
	DistributionConfigurationArn() *string
	// The name of the distribution configuration.
	// Experimental.
	DistributionConfigurationName() *string
}

// The jsii proxy for IDistributionConfiguration
type jsiiProxy_IDistributionConfiguration struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDistributionConfiguration) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
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

func (i *jsiiProxy_IDistributionConfiguration) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDistributionConfiguration) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistributionConfiguration) DistributionConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationName",
		&returns,
	)
	return returns
}

