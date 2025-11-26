package previewawsosismixins


// Options that specify the configuration of a persistent buffer.
//
// To configure how OpenSearch Ingestion encrypts this data, set the `EncryptionAtRestOptions` . For more information, see [Persistent buffering](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/osis-features-overview.html#persistent-buffering) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bufferOptionsProperty := &BufferOptionsProperty{
//   	PersistentBufferEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-bufferoptions.html
//
type CfnPipelinePropsMixin_BufferOptionsProperty struct {
	// Whether persistent buffering should be enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-bufferoptions.html#cfn-osis-pipeline-bufferoptions-persistentbufferenabled
	//
	PersistentBufferEnabled interface{} `field:"optional" json:"persistentBufferEnabled" yaml:"persistentBufferEnabled"`
}

