package previewawscodebuildevents


// Type definition for Additional-informationItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalInformationItem := &AdditionalInformationItem{
//   	DurationInSeconds: []*string{
//   		jsii.String("durationInSeconds"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	PhaseContext: []*string{
//   		jsii.String("phaseContext"),
//   	},
//   	PhaseStatus: []*string{
//   		jsii.String("phaseStatus"),
//   	},
//   	PhaseType: []*string{
//   		jsii.String("phaseType"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildStateChange_AdditionalInformationItem struct {
	// duration-in-seconds property.
	//
	// Specify an array of string values to match this event if the actual value of duration-in-seconds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DurationInSeconds *[]*string `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// end-time property.
	//
	// Specify an array of string values to match this event if the actual value of end-time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// phase-context property.
	//
	// Specify an array of string values to match this event if the actual value of phase-context is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseContext *[]*string `field:"optional" json:"phaseContext" yaml:"phaseContext"`
	// phase-status property.
	//
	// Specify an array of string values to match this event if the actual value of phase-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseStatus *[]*string `field:"optional" json:"phaseStatus" yaml:"phaseStatus"`
	// phase-type property.
	//
	// Specify an array of string values to match this event if the actual value of phase-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseType *[]*string `field:"optional" json:"phaseType" yaml:"phaseType"`
	// start-time property.
	//
	// Specify an array of string values to match this event if the actual value of start-time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
}

