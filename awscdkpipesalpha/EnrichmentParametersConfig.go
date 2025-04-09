package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes"
)

// The parameters required to set up enrichment on your pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   enrichmentParametersConfig := &EnrichmentParametersConfig{
//   	EnrichmentParameters: &PipeEnrichmentParametersProperty{
//   		HttpParameters: &PipeEnrichmentHttpParametersProperty{
//   			HeaderParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			PathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			QueryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		InputTemplate: jsii.String("inputTemplate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipeenrichmentparameters.html
//
// Experimental.
type EnrichmentParametersConfig struct {
	// The parameters for the enrichment target.
	// Experimental.
	EnrichmentParameters *awspipes.CfnPipe_PipeEnrichmentParametersProperty `field:"required" json:"enrichmentParameters" yaml:"enrichmentParameters"`
}

