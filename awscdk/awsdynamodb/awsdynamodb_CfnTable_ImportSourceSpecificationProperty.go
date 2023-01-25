package awsdynamodb


// Specifies the properties of data being imported from the S3 bucket source to the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   importSourceSpecificationProperty := &importSourceSpecificationProperty{
//   	inputFormat: jsii.String("inputFormat"),
//   	s3BucketSource: &s3BucketSourceProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		s3BucketOwner: jsii.String("s3BucketOwner"),
//   		s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//
//   	// the properties below are optional
//   	inputCompressionType: jsii.String("inputCompressionType"),
//   	inputFormatOptions: &inputFormatOptionsProperty{
//   		csv: &csvProperty{
//   			delimiter: jsii.String("delimiter"),
//   			headerList: []*string{
//   				jsii.String("headerList"),
//   			},
//   		},
//   	},
//   }
//
type CfnTable_ImportSourceSpecificationProperty struct {
	// The format of the source data.
	//
	// Valid values for `ImportFormat` are `CSV` , `DYNAMODB_JSON` or `ION` .
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// The S3 bucket that provides the source for the import.
	S3BucketSource interface{} `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// Type of compression to be used on the input coming from the imported table.
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// Additional properties that specify how the input is formatted,.
	InputFormatOptions interface{} `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

