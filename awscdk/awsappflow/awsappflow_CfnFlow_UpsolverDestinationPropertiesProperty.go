package awsappflow


// The properties that are applied when Upsolver is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upsolverDestinationPropertiesProperty := &upsolverDestinationPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   		prefixConfig: &prefixConfigProperty{
//   			pathPrefixHierarchy: []*string{
//   				jsii.String("pathPrefixHierarchy"),
//   			},
//   			prefixFormat: jsii.String("prefixFormat"),
//   			prefixType: jsii.String("prefixType"),
//   		},
//
//   		// the properties below are optional
//   		aggregationConfig: &aggregationConfigProperty{
//   			aggregationType: jsii.String("aggregationType"),
//   			targetFileSize: jsii.Number(123),
//   		},
//   		fileType: jsii.String("fileType"),
//   	},
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   }
//
type CfnFlow_UpsolverDestinationPropertiesProperty struct {
	// The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The configuration that determines how data is formatted when Upsolver is used as the flow destination.
	S3OutputFormatConfig interface{} `field:"required" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
	// The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

