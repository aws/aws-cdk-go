package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeEnrichmentParametersProperty := &pipeEnrichmentParametersProperty{
//   	httpParameters: &pipeEnrichmentHttpParametersProperty{
//   		headerParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		pathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		queryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	inputTemplate: jsii.String("inputTemplate"),
//   }
//
type CfnPipe_PipeEnrichmentParametersProperty struct {
	// `CfnPipe.PipeEnrichmentParametersProperty.HttpParameters`.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// `CfnPipe.PipeEnrichmentParametersProperty.InputTemplate`.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

