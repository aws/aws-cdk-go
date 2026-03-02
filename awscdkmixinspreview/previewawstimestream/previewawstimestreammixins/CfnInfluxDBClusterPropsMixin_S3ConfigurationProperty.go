package previewawstimestreammixins


// S3 configuration for sending logs to customer account from the InfluxDB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-s3configuration.html
//
type CfnInfluxDBClusterPropsMixin_S3ConfigurationProperty struct {
	// The bucket name for logs to be sent from the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-s3configuration.html#cfn-timestream-influxdbcluster-s3configuration-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies whether logging to customer specified bucket is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-s3configuration.html#cfn-timestream-influxdbcluster-s3configuration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

