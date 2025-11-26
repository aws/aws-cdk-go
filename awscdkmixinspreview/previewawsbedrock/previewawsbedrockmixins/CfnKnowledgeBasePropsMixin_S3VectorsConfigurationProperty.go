package previewawsbedrockmixins


// Contains the storage configuration of the knowledge base for S3 vectors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3VectorsConfigurationProperty := &S3VectorsConfigurationProperty{
//   	IndexArn: jsii.String("indexArn"),
//   	IndexName: jsii.String("indexName"),
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-s3vectorsconfiguration.html
//
type CfnKnowledgeBasePropsMixin_S3VectorsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the vector index used for the knowledge base.
	//
	// This ARN identifies the specific vector index resource within Amazon Bedrock.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-s3vectorsconfiguration.html#cfn-bedrock-knowledgebase-s3vectorsconfiguration-indexarn
	//
	IndexArn *string `field:"optional" json:"indexArn" yaml:"indexArn"`
	// The name of the vector index used for the knowledge base.
	//
	// This name identifies the vector index within the Amazon Bedrock service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-s3vectorsconfiguration.html#cfn-bedrock-knowledgebase-s3vectorsconfiguration-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The Amazon Resource Name (ARN) of the S3 bucket where vector embeddings are stored.
	//
	// This bucket contains the vector data used by the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-s3vectorsconfiguration.html#cfn-bedrock-knowledgebase-s3vectorsconfiguration-vectorbucketarn
	//
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
}

