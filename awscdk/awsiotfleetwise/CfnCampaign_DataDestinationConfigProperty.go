package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataDestinationConfigProperty := &DataDestinationConfigProperty{
//   	S3Config: &S3ConfigProperty{
//   		BucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		DataFormat: jsii.String("dataFormat"),
//   		Prefix: jsii.String("prefix"),
//   		StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   	},
//   	TimestreamConfig: &TimestreamConfigProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		TimestreamTableArn: jsii.String("timestreamTableArn"),
//   	},
//   }
//
type CfnCampaign_DataDestinationConfigProperty struct {
	// `CfnCampaign.DataDestinationConfigProperty.S3Config`.
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
	// `CfnCampaign.DataDestinationConfigProperty.TimestreamConfig`.
	TimestreamConfig interface{} `field:"optional" json:"timestreamConfig" yaml:"timestreamConfig"`
}

