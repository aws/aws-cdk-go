package awsdynamodb


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
	// `CfnTable.ImportSourceSpecificationProperty.InputFormat`.
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// `CfnTable.ImportSourceSpecificationProperty.S3BucketSource`.
	S3BucketSource interface{} `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// `CfnTable.ImportSourceSpecificationProperty.InputCompressionType`.
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// `CfnTable.ImportSourceSpecificationProperty.InputFormatOptions`.
	InputFormatOptions interface{} `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

