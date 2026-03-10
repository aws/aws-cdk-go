package previewawscodeguruprofilerevents


// Type definition for Recommendation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recommendation := &Recommendation{
//   	Description: &Description{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   	Name: &Name{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   	Reason: &Reason{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   	ResolutionSteps: &ResolutionSteps{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type CodeGuruProfilerRecommendationStateChange_Recommendation struct {
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *CodeGuruProfilerRecommendationStateChange_Description `field:"optional" json:"description" yaml:"description"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *CodeGuruProfilerRecommendationStateChange_Name `field:"optional" json:"name" yaml:"name"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *CodeGuruProfilerRecommendationStateChange_Reason `field:"optional" json:"reason" yaml:"reason"`
	// resolutionSteps property.
	//
	// Specify an array of string values to match this event if the actual value of resolutionSteps is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResolutionSteps *CodeGuruProfilerRecommendationStateChange_ResolutionSteps `field:"optional" json:"resolutionSteps" yaml:"resolutionSteps"`
}

