package awsappflow


// The properties that are applied when Upsolver is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   upsolverDestinationPropertiesProperty := &UpsolverDestinationPropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//   	S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
//   		PrefixConfig: &PrefixConfigProperty{
//   			PathPrefixHierarchy: []*string{
//   				jsii.String("pathPrefixHierarchy"),
//   			},
//   			PrefixFormat: jsii.String("prefixFormat"),
//   			PrefixType: jsii.String("prefixType"),
//   		},
//
//   		// the properties below are optional
//   		AggregationConfig: &AggregationConfigProperty{
//   			AggregationType: jsii.String("aggregationType"),
//   			TargetFileSize: jsii.Number(123),
//   		},
//   		FileType: jsii.String("fileType"),
//   	},
//
//   	// the properties below are optional
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html
//
type CfnFlow_UpsolverDestinationPropertiesProperty struct {
	// The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The configuration that determines how data is formatted when Upsolver is used as the flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-s3outputformatconfig
	//
	S3OutputFormatConfig interface{} `field:"required" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
	// The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

