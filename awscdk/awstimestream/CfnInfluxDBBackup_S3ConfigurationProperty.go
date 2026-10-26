package awstimestream


// Configuration for S3 bucket log delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-s3configuration.html
//
type CfnInfluxDBBackup_S3ConfigurationProperty struct {
	// The name of the S3 bucket to deliver logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-s3configuration.html#cfn-timestream-influxdbbackup-s3configuration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Indicates whether log delivery to the S3 bucket is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbbackup-s3configuration.html#cfn-timestream-influxdbbackup-s3configuration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

