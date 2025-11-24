package mixinsawss3


// The destination information for the S3 Metadata configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metadataDestinationProperty := &MetadataDestinationProperty{
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   	TableBucketType: jsii.String("tableBucketType"),
//   	TableNamespace: jsii.String("tableNamespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatadestination.html
//
type CfnBucketPropsMixin_MetadataDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the table bucket where the metadata configuration is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatadestination.html#cfn-s3-bucket-metadatadestination-tablebucketarn
	//
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The type of the table bucket where the metadata configuration is stored.
	//
	// The `aws` value indicates an AWS managed table bucket, and the `customer` value indicates a customer-managed table bucket. V2 metadata configurations are stored in AWS managed table buckets, and V1 metadata configurations are stored in customer-managed table buckets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatadestination.html#cfn-s3-bucket-metadatadestination-tablebuckettype
	//
	TableBucketType *string `field:"optional" json:"tableBucketType" yaml:"tableBucketType"`
	// The namespace in the table bucket where the metadata tables for a metadata configuration are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatadestination.html#cfn-s3-bucket-metadatadestination-tablenamespace
	//
	TableNamespace *string `field:"optional" json:"tableNamespace" yaml:"tableNamespace"`
}

