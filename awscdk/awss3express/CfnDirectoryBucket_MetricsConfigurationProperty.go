package awss3express


// Specifies a metrics configuration for the CloudWatch request metrics from an Amazon S3 Express bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigurationProperty := &MetricsConfigurationProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html
//
type CfnDirectoryBucket_MetricsConfigurationProperty struct {
	// The ID used to identify the metrics configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The access point ARN used when evaluating a metrics filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The prefix used when evaluating a metrics filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

