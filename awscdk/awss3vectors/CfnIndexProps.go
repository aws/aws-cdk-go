package awss3vectors


// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &CfnIndexProps{
//   	DataType: jsii.String("dataType"),
//   	Dimension: jsii.Number(123),
//   	DistanceMetric: jsii.String("distanceMetric"),
//
//   	// the properties below are optional
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
type CfnIndexProps struct {
	// The data type of the vectors to be inserted into the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-datatype
	//
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The dimensions of the vectors to be inserted into the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-dimension
	//
	Dimension *float64 `field:"required" json:"dimension" yaml:"dimension"`
	// The distance metric to be used for similarity search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-distancemetric
	//
	DistanceMetric *string `field:"required" json:"distanceMetric" yaml:"distanceMetric"`
	// The name of the vector index to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The metadata configuration for the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The Amazon Resource Name (ARN) of the vector bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-vectorbucketarn
	//
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The name of the vector bucket that contains the vector index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3vectors-index.html#cfn-s3vectors-index-vectorbucketname
	//
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

