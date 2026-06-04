package awsbedrockagentcore


// Source of initial examples.
//
// Provide either inline examples or an S3 URI pointing to a JSONL file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var examples interface{}
//
//   dataSourceTypeProperty := &DataSourceTypeProperty{
//   	InlineExamples: &InlineExamplesSourceProperty{
//   		Examples: []interface{}{
//   			examples,
//   		},
//   	},
//   	S3Source: &S3SourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-datasourcetype.html
//
type CfnDatasetPropsMixin_DataSourceTypeProperty struct {
	// Inline examples provided directly in the request body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-datasourcetype.html#cfn-bedrockagentcore-dataset-datasourcetype-inlineexamples
	//
	InlineExamples interface{} `field:"optional" json:"inlineExamples" yaml:"inlineExamples"`
	// S3 location of a JSONL file containing dataset examples.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-datasourcetype.html#cfn-bedrockagentcore-dataset-datasourcetype-s3source
	//
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
}

