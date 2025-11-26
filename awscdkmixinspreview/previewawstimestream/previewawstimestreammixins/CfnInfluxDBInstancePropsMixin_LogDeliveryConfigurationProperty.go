package previewawstimestreammixins


// Configuration for sending InfluxDB engine logs to a specified S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logDeliveryConfigurationProperty := &LogDeliveryConfigurationProperty{
//   	S3Configuration: &S3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbinstance-logdeliveryconfiguration.html
//
type CfnInfluxDBInstancePropsMixin_LogDeliveryConfigurationProperty struct {
	// Configuration for S3 bucket log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbinstance-logdeliveryconfiguration.html#cfn-timestream-influxdbinstance-logdeliveryconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

