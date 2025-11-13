package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Infrastructure Configuration.
// Experimental.
type IInfrastructureConfiguration interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the infrastructure configuration.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the infrastructure configuration.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the infrastructure configuration.
	// Experimental.
	InfrastructureConfigurationArn() *string
	// The name of the infrastructure configuration.
	// Experimental.
	InfrastructureConfigurationName() *string
}

// The jsii proxy for IInfrastructureConfiguration
type jsiiProxy_IInfrastructureConfiguration struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInfrastructureConfiguration) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IInfrastructureConfiguration) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (j *jsiiProxy_IInfrastructureConfiguration) InfrastructureConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInfrastructureConfiguration) InfrastructureConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationName",
		&returns,
	)
	return returns
}

