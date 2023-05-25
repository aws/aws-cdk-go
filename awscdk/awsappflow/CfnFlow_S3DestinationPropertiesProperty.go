package awsappflow


// The properties that are applied when Amazon S3 is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationPropertiesProperty := &S3DestinationPropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	S3OutputFormatConfig: &S3OutputFormatConfigProperty{
//   		AggregationConfig: &AggregationConfigProperty{
//   			AggregationType: jsii.String("aggregationType"),
//   			TargetFileSize: jsii.Number(123),
//   		},
//   		FileType: jsii.String("fileType"),
//   		PrefixConfig: &PrefixConfigProperty{
//   			PathPrefixHierarchy: []*string{
//   				jsii.String("pathPrefixHierarchy"),
//   			},
//   			PrefixFormat: jsii.String("prefixFormat"),
//   			PrefixType: jsii.String("prefixType"),
//   		},
//   		PreserveSourceDataTyping: jsii.Boolean(false),
//   	},
//   }
//
type CfnFlow_S3DestinationPropertiesProperty struct {
	// The Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination.
	S3OutputFormatConfig interface{} `field:"optional" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}

