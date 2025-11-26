package previewawss3vectorsmixins


// Properties for CfnIndexPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIndexMixinProps := &CfnIndexMixinProps{
//   	DataType: jsii.String("dataType"),
//   	Dimension: jsii.Number(123),
//   	DistanceMetric: jsii.String("distanceMetric"),
//   	IndexName: jsii.String("indexName"),
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		NonFilterableMetadataKeys: []*string{
//   			jsii.String("nonFilterableMetadataKeys"),
//   		},
//   	},
//   	VectorBucketArn: jsii.String("vectorBucketArn"),
//   	VectorBucketName: jsii.String("vectorBucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html
//
type CfnIndexMixinProps struct {
	// The data type of the vectors to be inserted into the vector index.
	//
	// Currently, only `float32` is supported, which represents 32-bit floating-point numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The dimensions of the vectors to be inserted into the vector index.
	//
	// This value must be between 1 and 4096, inclusive. All vectors stored in the index must have the same number of dimensions.
	//
	// The dimension value affects the storage requirements and search performance. Higher dimensions require more storage space and may impact search latency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-dimension
	//
	Dimension *float64 `field:"optional" json:"dimension" yaml:"dimension"`
	// The distance metric to be used for similarity search. Valid values are:.
	//
	// - `cosine` - Measures the cosine of the angle between two vectors.
	// - `euclidean` - Measures the straight-line distance between two points in multi-dimensional space. Lower values indicate greater similarity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-distancemetric
	//
	DistanceMetric *string `field:"optional" json:"distanceMetric" yaml:"distanceMetric"`
	// The name of the vector index to create.
	//
	// The index name must be between 3 and 63 characters long and can contain only lowercase letters, numbers, hyphens (-), and dots (.). The index name must be unique within the vector bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the index name.
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The metadata configuration for the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The Amazon Resource Name (ARN) of the vector bucket that contains the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-vectorbucketarn
	//
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The name of the vector bucket that contains the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-vectorbucketname
	//
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

