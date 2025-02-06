package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Base properties for calling an REST API Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi restApi
//
//   callApiGatewayRestApiEndpointOptions := &CallApiGatewayRestApiEndpointOptions{
//   	Api: restApi,
//   	StageName: jsii.String("stageName"),
//   }
//
type CallApiGatewayRestApiEndpointOptions struct {
	// API to call.
	Api awsapigateway.IRestApi `field:"required" json:"api" yaml:"api"`
	// Name of the stage where the API is deployed to in API Gateway.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

