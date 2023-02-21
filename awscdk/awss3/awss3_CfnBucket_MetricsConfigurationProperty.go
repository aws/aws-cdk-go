package awss3


// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
//
// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For examples, see [AWS::S3::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples) . For more information, see [PUT Bucket metrics](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigurationProperty := &metricsConfigurationProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	accessPointArn: jsii.String("accessPointArn"),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_MetricsConfigurationProperty struct {
	// The ID used to identify the metrics configuration.
	//
	// This can be any value you choose that helps you identify your metrics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The access point that was used while performing operations on the object.
	//
	// The metrics configuration only includes objects that meet the filter's criteria.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The prefix that an object must have to be included in the metrics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a list of tag filters to use as a metrics configuration filter.
	//
	// The metrics configuration includes only objects that meet the filter's criteria.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

