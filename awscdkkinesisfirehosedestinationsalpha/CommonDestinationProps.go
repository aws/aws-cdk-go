// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2"
)

// Generic properties for defining a delivery stream destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import kinesisfirehose_destinations_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var compression compression
//   var dataProcessor iDataProcessor
//   var key key
//   var logGroup logGroup
//   var role role
//   var size size
//
//   commonDestinationProps := &CommonDestinationProps{
//   	Logging: jsii.Boolean(false),
//   	LogGroup: logGroup,
//   	Processor: dataProcessor,
//   	Role: role,
//   	S3Backup: &DestinationS3BackupProps{
//   		Bucket: bucket,
//   		BufferingInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   		BufferingSize: size,
//   		Compression: compression,
//   		DataOutputPrefix: jsii.String("dataOutputPrefix"),
//   		EncryptionKey: key,
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Logging: jsii.Boolean(false),
//   		LogGroup: logGroup,
//   		Mode: kinesisfirehose_destinations_alpha.BackupMode_ALL,
//   	},
//   }
//
// Experimental.
type CommonDestinationProps struct {
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Experimental.
	Processor awscdkkinesisfirehosealpha.IDataProcessor `field:"optional" json:"processor" yaml:"processor"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The configuration for backing up source records to S3.
	// Experimental.
	S3Backup *DestinationS3BackupProps `field:"optional" json:"s3Backup" yaml:"s3Backup"`
}

