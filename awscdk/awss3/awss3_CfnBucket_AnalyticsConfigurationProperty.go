package awss3


// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyticsConfigurationProperty := &analyticsConfigurationProperty{
//   	id: jsii.String("id"),
//   	storageClassAnalysis: &storageClassAnalysisProperty{
//   		dataExport: &dataExportProperty{
//   			destination: &destinationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				format: jsii.String("format"),
//
//   				// the properties below are optional
//   				bucketAccountId: jsii.String("bucketAccountId"),
//   				prefix: jsii.String("prefix"),
//   			},
//   			outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_AnalyticsConfigurationProperty struct {
	// The ID that identifies the analytics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes.
	StorageClassAnalysis interface{} `field:"required" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
	// The prefix that an object must have to be included in the analytics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The tags to use when evaluating an analytics filter.
	//
	// The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

