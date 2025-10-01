package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Base properties for calling an HTTP API Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   callApiGatewayHttpApiEndpointOptions := &CallApiGatewayHttpApiEndpointOptions{
//   	ApiId: jsii.String("apiId"),
//   	ApiStack: stack,
//
//   	// the properties below are optional
//   	StageName: jsii.String("stageName"),
//   }
//
type CallApiGatewayHttpApiEndpointOptions struct {
	// The Id of the API to call.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The Stack in which the API is defined.
	ApiStack awscdk.Stack `field:"required" json:"apiStack" yaml:"apiStack"`
	// Name of the stage where the API is deployed to in API Gateway.
	// Default: '$default'.
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

