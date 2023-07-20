package awsdynamodb


// Specifies the properties of data being imported from the S3 bucket source to the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   importSourceSpecificationProperty := &ImportSourceSpecificationProperty{
//   	InputFormat: jsii.String("inputFormat"),
//   	S3BucketSource: &S3BucketSourceProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//
//   	// the properties below are optional
//   	InputCompressionType: jsii.String("inputCompressionType"),
//   	InputFormatOptions: &InputFormatOptionsProperty{
//   		Csv: &CsvProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			HeaderList: []*string{
//   				jsii.String("headerList"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-importsourcespecification.html
//
type CfnTable_ImportSourceSpecificationProperty struct {
	// The format of the source data.
	//
	// Valid values for `ImportFormat` are `CSV` , `DYNAMODB_JSON` or `ION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-importsourcespecification.html#cfn-dynamodb-table-importsourcespecification-inputformat
	//
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// The S3 bucket that provides the source for the import.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-importsourcespecification.html#cfn-dynamodb-table-importsourcespecification-s3bucketsource
	//
	S3BucketSource interface{} `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// Type of compression to be used on the input coming from the imported table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-importsourcespecification.html#cfn-dynamodb-table-importsourcespecification-inputcompressiontype
	//
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// Additional properties that specify how the input is formatted,.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-importsourcespecification.html#cfn-dynamodb-table-importsourcespecification-inputformatoptions
	//
	InputFormatOptions interface{} `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

