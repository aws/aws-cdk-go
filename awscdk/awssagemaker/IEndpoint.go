package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
)

// The interface for a SageMaker Endpoint resource.
type IEndpoint interface {
	awscdk.IResource
	// Permits an IAM principal to invoke this endpoint.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the endpoint.
	EndpointArn() *string
	// The name of the endpoint.
	EndpointName() *string
}

// The jsii proxy for IEndpoint
type jsiiProxy_IEndpoint struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEndpoint) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpoint) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

