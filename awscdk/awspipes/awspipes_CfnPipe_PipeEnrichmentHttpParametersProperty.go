package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeEnrichmentHttpParametersProperty := &pipeEnrichmentHttpParametersProperty{
//   	headerParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	pathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   }
//
type CfnPipe_PipeEnrichmentHttpParametersProperty struct {
	// `CfnPipe.PipeEnrichmentHttpParametersProperty.HeaderParameters`.
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// `CfnPipe.PipeEnrichmentHttpParametersProperty.PathParameterValues`.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// `CfnPipe.PipeEnrichmentHttpParametersProperty.QueryStringParameters`.
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

