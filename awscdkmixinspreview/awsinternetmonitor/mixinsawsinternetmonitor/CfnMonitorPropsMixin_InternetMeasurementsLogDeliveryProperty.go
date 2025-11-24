package mixinsawsinternetmonitor


// Publish internet measurements to an Amazon S3 bucket in addition to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   internetMeasurementsLogDeliveryProperty := &InternetMeasurementsLogDeliveryProperty{
//   	S3Config: &S3ConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		LogDeliveryStatus: jsii.String("logDeliveryStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html
//
type CfnMonitorPropsMixin_InternetMeasurementsLogDeliveryProperty struct {
	// The configuration for publishing Amazon CloudWatch Internet Monitor internet measurements to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html#cfn-internetmonitor-monitor-internetmeasurementslogdelivery-s3config
	//
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
}

