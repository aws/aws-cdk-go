package awscdkpipesenrichmentsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Properties for a ApiDestinationEnrichment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import pipes_enrichments_alpha "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha"
//
//   var inputTransformation inputTransformation
//
//   apiDestinationEnrichmentProps := &ApiDestinationEnrichmentProps{
//   	HeaderParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	InputTransformation: inputTransformation,
//   	PathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   }
//
// Experimental.
type ApiDestinationEnrichmentProps struct {
	// The headers that need to be sent as part of request invoking the EventBridge ApiDestination.
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
	// The path parameter values used to populate the EventBridge API destination path wildcards ("*").
	// Default: - none.
	//
	// Experimental.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the EventBridge API destination.
	// Default: - none.
	//
	// Experimental.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

