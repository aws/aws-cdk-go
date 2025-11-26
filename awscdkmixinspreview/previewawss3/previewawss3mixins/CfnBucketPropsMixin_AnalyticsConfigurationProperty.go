package previewawss3mixins


// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analyticsConfigurationProperty := &AnalyticsConfigurationProperty{
//   	Id: jsii.String("id"),
//   	Prefix: jsii.String("prefix"),
//   	StorageClassAnalysis: &StorageClassAnalysisProperty{
//   		DataExport: &DataExportProperty{
//   			Destination: &DestinationProperty{
//   				BucketAccountId: jsii.String("bucketAccountId"),
//   				BucketArn: jsii.String("bucketArn"),
//   				Format: jsii.String("format"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		},
//   	},
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
type CfnBucketPropsMixin_AnalyticsConfigurationProperty struct {
	// The ID that identifies the analytics configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The prefix that an object must have to be included in the analytics results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-storageclassanalysis
	//
	StorageClassAnalysis interface{} `field:"optional" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
	// The tags to use when evaluating an analytics filter.
	//
	// The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-tagfilters
	//
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

