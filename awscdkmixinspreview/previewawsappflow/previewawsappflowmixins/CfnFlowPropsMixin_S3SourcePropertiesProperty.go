package previewawsappflowmixins


// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3SourcePropertiesProperty := &S3SourcePropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	S3InputFormatConfig: &S3InputFormatConfigProperty{
//   		S3InputFileType: jsii.String("s3InputFileType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3sourceproperties.html
//
type CfnFlowPropsMixin_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3sourceproperties.html#cfn-appflow-flow-s3sourceproperties-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3sourceproperties.html#cfn-appflow-flow-s3sourceproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// When you use Amazon S3 as the source, the configuration format that you provide the flow input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3sourceproperties.html#cfn-appflow-flow-s3sourceproperties-s3inputformatconfig
	//
	S3InputFormatConfig interface{} `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

