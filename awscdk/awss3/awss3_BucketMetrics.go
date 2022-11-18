package awss3


// Specifies a metrics configuration for the CloudWatch request metrics from an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagFilters interface{}
//
//   bucketMetrics := &bucketMetrics{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   	tagFilters: map[string]interface{}{
//   		"tagFiltersKey": tagFilters,
//   	},
//   }
//
type BucketMetrics struct {
	// The ID used to identify the metrics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The prefix that an object must have to be included in the metrics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a list of tag filters to use as a metrics configuration filter.
	//
	// The metrics configuration includes only objects that meet the filter's criteria.
	TagFilters *map[string]interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

