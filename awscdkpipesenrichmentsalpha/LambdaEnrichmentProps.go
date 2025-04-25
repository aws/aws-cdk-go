package awscdkpipesenrichmentsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Properties for a LambdaEnrichment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import pipes_enrichments_alpha "github.com/aws/aws-cdk-go/awscdkpipesenrichmentsalpha"
//
//   var inputTransformation inputTransformation
//
//   lambdaEnrichmentProps := &LambdaEnrichmentProps{
//   	InputTransformation: inputTransformation,
//   }
//
// Experimental.
type LambdaEnrichmentProps struct {
	// The input transformation for the enrichment.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-input-transformation.html
	//
	// Default: - None.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.InputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
}

