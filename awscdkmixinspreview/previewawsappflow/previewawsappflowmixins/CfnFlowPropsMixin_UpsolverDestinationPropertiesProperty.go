package previewawsappflowmixins


// The properties that are applied when Upsolver is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   upsolverDestinationPropertiesProperty := &UpsolverDestinationPropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html
//
type CfnFlowPropsMixin_UpsolverDestinationPropertiesProperty struct {
	// The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the transferred data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The configuration that determines how data is formatted when Upsolver is used as the flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-upsolverdestinationproperties.html#cfn-appflow-flow-upsolverdestinationproperties-s3outputformatconfig
	//
	S3OutputFormatConfig interface{} `field:"optional" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}

