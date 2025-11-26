package previewawsguarddutyevents


// Type definition for HighestSeverityThreatDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   highestSeverityThreatDetails := &HighestSeverityThreatDetails{
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	ThreatName: []*string{
//   		jsii.String("threatName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_HighestSeverityThreatDetails struct {
	// count property.
	//
	// Specify an array of string values to match this event if the actual value of count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// threatName property.
	//
	// Specify an array of string values to match this event if the actual value of threatName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatName *[]*string `field:"optional" json:"threatName" yaml:"threatName"`
}

