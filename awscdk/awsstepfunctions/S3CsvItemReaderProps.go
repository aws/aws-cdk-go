package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for configuring an Item Reader that iterates over items in a CSV file in S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var csvHeaders csvHeaders
//
//   s3CsvItemReaderProps := &S3CsvItemReaderProps{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Bucket: bucket,
//   	BucketNamePath: jsii.String("bucketNamePath"),
//   	CsvDelimiter: awscdk.Aws_stepfunctions.CsvDelimiter_COMMA,
//   	CsvHeaders: csvHeaders,
//   	MaxItems: jsii.Number(123),
//   }
//
type S3CsvItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	// See: bucketNamePath.
	//
	// Default: - S3 bucket will be determined from.
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// S3 bucket name containing objects to iterate over or a file with a list to iterate over, as JsonPath.
	// See: bucket.
	//
	// Default: - S3 bucket will be determined from.
	//
	BucketNamePath *string `field:"optional" json:"bucketNamePath" yaml:"bucketNamePath"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Key of file stored in S3 bucket containing an array to iterate over.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Delimiter used in a CSV file.
	// Default: undefined - Default setting is COMMA.
	//
	CsvDelimiter CsvDelimiter `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// CSV file header configuration.
	// Default: - CsvHeaders with CsvHeadersLocation.FIRST_ROW
	//
	CsvHeaders CsvHeaders `field:"optional" json:"csvHeaders" yaml:"csvHeaders"`
}

