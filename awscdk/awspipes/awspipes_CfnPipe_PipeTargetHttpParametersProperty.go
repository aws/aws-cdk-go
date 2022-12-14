package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetHttpParametersProperty := &pipeTargetHttpParametersProperty{
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
type CfnPipe_PipeTargetHttpParametersProperty struct {
	// `CfnPipe.PipeTargetHttpParametersProperty.HeaderParameters`.
	HeaderParameters interface{} `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// `CfnPipe.PipeTargetHttpParametersProperty.PathParameterValues`.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// `CfnPipe.PipeTargetHttpParametersProperty.QueryStringParameters`.
	QueryStringParameters interface{} `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

