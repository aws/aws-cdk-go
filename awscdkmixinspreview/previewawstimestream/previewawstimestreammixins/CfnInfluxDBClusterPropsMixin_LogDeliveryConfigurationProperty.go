package previewawstimestreammixins


// Configuration for sending logs to customer account from the InfluxDB cluster.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-logdeliveryconfiguration.html
//
type CfnInfluxDBClusterPropsMixin_LogDeliveryConfigurationProperty struct {
	// S3 configuration for sending logs to customer account from the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-influxdbcluster-logdeliveryconfiguration.html#cfn-timestream-influxdbcluster-logdeliveryconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

