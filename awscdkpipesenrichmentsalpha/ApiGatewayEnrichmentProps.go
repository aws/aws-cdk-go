package awscdkpipesenrichmentsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Properties for a ApiGatewayEnrichment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import pipes_enrichments_alpha "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha"
//
//   var inputTransformation InputTransformation
//
//   apiGatewayEnrichmentProps := &ApiGatewayEnrichmentProps{
//   	HeaderParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	InputTransformation: inputTransformation,
//   	Method: jsii.String("method"),
//   	Path: jsii.String("path"),
//   	PathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   	Stage: jsii.String("stage"),
//   }
//
// Experimental.
type ApiGatewayEnrichmentProps struct {
	// The headers that need to be sent as part of request invoking the API Gateway REST API.
	// Default: - none.
	//
	// Experimental.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The input transformation for the enrichment.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-input-transformation.html
	//
	// Default: - None.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.InputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// The method for API Gateway resource.
	// Default: '*' - ANY.
	//
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The path for the API Gateway resource.
	// Default: '/'.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The path parameter values used to populate the API Gateway REST API path wildcards ("*").
	// Default: - none.
	//
	// Experimental.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the EventBridge API destination.
	// Default: - none.
	//
	// Experimental.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
	// The deployment stage for the API Gateway resource.
	// Default: - the value of `deploymentStage.stageName` of target API Gateway resource.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

