package awslookoutmetrics


// Contains information about the configuration of the S3 bucket that contains source files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceConfigProperty := &s3SourceConfigProperty{
//   	fileFormatDescriptor: &fileFormatDescriptorProperty{
//   		csvFormatDescriptor: &csvFormatDescriptorProperty{
//   			charset: jsii.String("charset"),
//   			containsHeader: jsii.Boolean(false),
//   			delimiter: jsii.String("delimiter"),
//   			fileCompression: jsii.String("fileCompression"),
//   			headerList: []*string{
//   				jsii.String("headerList"),
//   			},
//   			quoteSymbol: jsii.String("quoteSymbol"),
//   		},
//   		jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   			charset: jsii.String("charset"),
//   			fileCompression: jsii.String("fileCompression"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	historicalDataPathList: []*string{
//   		jsii.String("historicalDataPathList"),
//   	},
//   	templatedPathList: []*string{
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

