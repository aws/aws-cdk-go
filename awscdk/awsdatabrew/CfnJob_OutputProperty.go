package awsdatabrew


// Represents options that specify how and where in Amazon S3 DataBrew writes the output generated by recipe jobs or profile jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &OutputProperty{
//   	Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	CompressionFormat: jsii.String("compressionFormat"),
//   	Format: jsii.String("format"),
//   	FormatOptions: &OutputFormatOptionsProperty{
//   		Csv: &CsvOutputOptionsProperty{
//   			Delimiter: jsii.String("delimiter"),
//   		},
//   	},
//   	MaxOutputFiles: jsii.Number(123),
//   	Overwrite: jsii.Boolean(false),
//   	PartitionColumns: []*string{
//   		jsii.String("partitionColumns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html
//
type CfnJob_OutputProperty struct {
	// The location in Amazon S3 where the job writes its output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-location
	//
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The compression algorithm used to compress the output text of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-compressionformat
	//
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// The data format of the output of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Represents options that define how DataBrew formats job output files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// The maximum number of files to be generated by the job and written to the output folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-maxoutputfiles
	//
	MaxOutputFiles *float64 `field:"optional" json:"maxOutputFiles" yaml:"maxOutputFiles"`
	// A value that, if true, means that any data in the location specified for output is overwritten with new output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-overwrite
	//
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// The names of one or more partition columns for the output of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-partitioncolumns
	//
	PartitionColumns *[]*string `field:"optional" json:"partitionColumns" yaml:"partitionColumns"`
}

