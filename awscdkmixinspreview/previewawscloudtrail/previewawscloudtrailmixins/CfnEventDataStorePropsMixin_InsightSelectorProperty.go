package previewawscloudtrailmixins


// A JSON string that contains a list of Insights types that are logged on an event data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   insightSelectorProperty := &InsightSelectorProperty{
//   	InsightType: jsii.String("insightType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-insightselector.html
//
type CfnEventDataStorePropsMixin_InsightSelectorProperty struct {
	// The type of Insights events to log on an event data store. `ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight types.
	//
	// The `ApiCallRateInsight` Insights type analyzes write-only management API calls that are aggregated per minute against a baseline API call volume.
	//
	// The `ApiErrorRateInsight` Insights type analyzes management API calls that result in error codes. The error is shown if the API call is unsuccessful.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-insightselector.html#cfn-cloudtrail-eventdatastore-insightselector-insighttype
	//
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

