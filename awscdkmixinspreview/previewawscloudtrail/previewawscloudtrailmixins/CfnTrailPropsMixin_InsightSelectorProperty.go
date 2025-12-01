package previewawscloudtrailmixins


// A JSON string that contains a list of Insights types that are logged on a trail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   insightSelectorProperty := &InsightSelectorProperty{
//   	EventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	InsightType: jsii.String("insightType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-insightselector.html
//
type CfnTrailPropsMixin_InsightSelectorProperty struct {
	// Select the event category on which Insights should be enabled.
	//
	// - If EventCategories is not provided, the specified Insights types are enabled on management API calls by default.
	// - If EventCategories is provided, the given event categories will overwrite the existing ones. For example, if a trail already has Insights enabled on management events, and then a PutInsightSelectors request is made with only data events specified in EventCategories, Insights on management events will be disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-insightselector.html#cfn-cloudtrail-trail-insightselector-eventcategories
	//
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// The type of Insights events to log on a trail. `ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight types.
	//
	// The `ApiCallRateInsight` Insights type analyzes write-only management API calls that are aggregated per minute against a baseline API call volume.
	//
	// The `ApiErrorRateInsight` Insights type analyzes management API calls that result in error codes. The error is shown if the API call is unsuccessful.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-insightselector.html#cfn-cloudtrail-trail-insightselector-insighttype
	//
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

