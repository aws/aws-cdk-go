package awslookoutmetrics


// Contains information about the configuration of the S3 bucket that contains source files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceConfigProperty := &S3SourceConfigProperty{
//   	FileFormatDescriptor: &FileFormatDescriptorProperty{
//   		CsvFormatDescriptor: &CsvFormatDescriptorProperty{
//   			Charset: jsii.String("charset"),
//   			ContainsHeader: jsii.Boolean(false),
//   			Delimiter: jsii.String("delimiter"),
//   			FileCompression: jsii.String("fileCompression"),
//   			HeaderList: []*string{
//   				jsii.String("headerList"),
//   			},
//   			QuoteSymbol: jsii.String("quoteSymbol"),
//   		},
//   		JsonFormatDescriptor: &JsonFormatDescriptorProperty{
//   			Charset: jsii.String("charset"),
//   			FileCompression: jsii.String("fileCompression"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	HistoricalDataPathList: []*string{
//   		jsii.String("historicalDataPathList"),
//   	},
//   	TemplatedPathList: []*string{
//   		jsii.String("templatedPathList"),
//   	},
//   }
//
type CfnAnomalyDetector_S3SourceConfigProperty struct {
	// Contains information about a source file's formatting.
	FileFormatDescriptor interface{} `field:"required" json:"fileFormatDescriptor" yaml:"fileFormatDescriptor"`
	// The ARN of an IAM role that has read and write access permissions to the source S3 bucket.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A list of paths to the historical data files.
	HistoricalDataPathList *[]*string `field:"optional" json:"historicalDataPathList" yaml:"historicalDataPathList"`
	// A list of templated paths to the source files.
	TemplatedPathList *[]*string `field:"optional" json:"templatedPathList" yaml:"templatedPathList"`
}

