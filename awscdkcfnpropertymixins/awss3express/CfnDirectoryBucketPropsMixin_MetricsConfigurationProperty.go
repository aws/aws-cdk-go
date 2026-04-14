package awss3express


// Specifies a metrics configuration for the CloudWatch request metrics from an Amazon S3 Express bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricsConfigurationProperty := &MetricsConfigurationProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	Id: jsii.String("id"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html
//
type CfnDirectoryBucketPropsMixin_MetricsConfigurationProperty struct {
	// The access point ARN used when evaluating a metrics filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID used to identify the metrics configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The prefix used when evaluating a metrics filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-metricsconfiguration.html#cfn-s3express-directorybucket-metricsconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

