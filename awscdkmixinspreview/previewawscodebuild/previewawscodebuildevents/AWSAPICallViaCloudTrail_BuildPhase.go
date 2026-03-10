package previewawscodebuildevents


// Type definition for BuildPhase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   buildPhase := &BuildPhase{
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
type AWSAPICallViaCloudTrail_BuildPhase struct {
	// durationInSeconds property.
	//
	// Specify an array of string values to match this event if the actual value of durationInSeconds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DurationInSeconds *[]*string `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// phaseContext property.
	//
	// Specify an array of string values to match this event if the actual value of phaseContext is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseContext *[]*string `field:"optional" json:"phaseContext" yaml:"phaseContext"`
	// phaseStatus property.
	//
	// Specify an array of string values to match this event if the actual value of phaseStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseStatus *[]*string `field:"optional" json:"phaseStatus" yaml:"phaseStatus"`
	// phaseType property.
	//
	// Specify an array of string values to match this event if the actual value of phaseType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PhaseType *[]*string `field:"optional" json:"phaseType" yaml:"phaseType"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
}

