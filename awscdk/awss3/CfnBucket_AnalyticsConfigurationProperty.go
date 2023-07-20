package awss3


// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyticsConfigurationProperty := &AnalyticsConfigurationProperty{
//   	Id: jsii.String("id"),
//   	StorageClassAnalysis: &StorageClassAnalysisProperty{
//   		DataExport: &DataExportProperty{
//   			Destination: &DestinationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				Format: jsii.String("format"),
//
//   				// the properties below are optional
//   				BucketAccountId: jsii.String("bucketAccountId"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   	TagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html
//
type CfnBucket_AnalyticsConfigurationProperty struct {
	// The ID that identifies the analytics configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-storageclassanalysis
	//
	StorageClassAnalysis interface{} `field:"required" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
	// The prefix that an object must have to be included in the analytics results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The tags to use when evaluating an analytics filter.
	//
	// The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-tagfilters
	//
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

