package previewawsguarddutyevents


// Type definition for ThreatsDetectedItemCount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   threatsDetectedItemCount := &ThreatsDetectedItemCount{
//   	Files: []*string{
//   		jsii.String("files"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ThreatsDetectedItemCount struct {
	// files property.
	//
	// Specify an array of string values to match this event if the actual value of files is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
}

