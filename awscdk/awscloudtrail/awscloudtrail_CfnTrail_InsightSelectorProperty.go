package awscloudtrail


// A JSON string that contains a list of insight types that are logged on a trail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightSelectorProperty := &insightSelectorProperty{
//   	insightType: jsii.String("insightType"),
//   }
//
type CfnTrail_InsightSelectorProperty struct {
	// The type of insights to log on a trail.
	//
	// `ApiCallRateInsight` and `ApiErrorRateInsight` are valid insight types.
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

