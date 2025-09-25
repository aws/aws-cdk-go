package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
)

// Represents an ApiGatewayV2 ApiMapping resource.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html
//
type IApiMapping interface {
	awscdk.IResource
	// ID of the api mapping.
	ApiMappingId() *string
}

// The jsii proxy for IApiMapping
type jsiiProxy_IApiMapping struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiMapping) ApiMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingId",
		&returns,
	)
	return returns
}

