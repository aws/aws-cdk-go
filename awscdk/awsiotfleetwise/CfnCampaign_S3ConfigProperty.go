package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &S3ConfigProperty{
//   	BucketArn: jsii.String("bucketArn"),
//
//   	// the properties below are optional
//   	DataFormat: jsii.String("dataFormat"),
//   	Prefix: jsii.String("prefix"),
//   	StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   }
//
type CfnCampaign_S3ConfigProperty struct {
	// `CfnCampaign.S3ConfigProperty.BucketArn`.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// `CfnCampaign.S3ConfigProperty.DataFormat`.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// `CfnCampaign.S3ConfigProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// `CfnCampaign.S3ConfigProperty.StorageCompressionFormat`.
	StorageCompressionFormat *string `field:"optional" json:"storageCompressionFormat" yaml:"storageCompressionFormat"`
}

