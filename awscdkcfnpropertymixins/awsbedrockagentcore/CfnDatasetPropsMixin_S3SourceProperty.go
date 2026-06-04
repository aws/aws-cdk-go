package awsbedrockagentcore


// S3 location of a JSONL file containing dataset examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3SourceProperty := &S3SourceProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-s3source.html
//
type CfnDatasetPropsMixin_S3SourceProperty struct {
	// S3 URI of the JSONL file (e.g. s3://my-bucket/path/to/examples.jsonl).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-s3source.html#cfn-bedrockagentcore-dataset-s3source-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

