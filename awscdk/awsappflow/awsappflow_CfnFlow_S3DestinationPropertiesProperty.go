package awsappflow


// The properties that are applied when Amazon S3 is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationPropertiesProperty := &s3DestinationPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   		aggregationConfig: &aggregationConfigProperty{
//   			aggregationType: jsii.String("aggregationType"),
//   			targetFileSize: jsii.Number(123),
//   		},
//   		fileType: jsii.String("fileType"),
//   		prefixConfig: &prefixConfigProperty{
//   			pathPrefixHierarchy: []*string{
//   				jsii.String("pathPrefixHierarchy"),
//   			},
//   			prefixFormat: jsii.String("prefixFormat"),
//   			prefixType: jsii.String("prefixType"),
//   		},
//   		preserveSourceDataTyping: jsii.Boolean(false),
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

