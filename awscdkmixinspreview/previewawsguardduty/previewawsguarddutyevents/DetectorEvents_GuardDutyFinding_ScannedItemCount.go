package previewawsguarddutyevents


// Type definition for ScannedItemCount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scannedItemCount := &ScannedItemCount{
//   	Files: []*string{
//   		jsii.String("files"),
//   	},
//   	TotalGb: []*string{
//   		jsii.String("totalGb"),
//   	},
//   	Volumes: []*string{
//   		jsii.String("volumes"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ScannedItemCount struct {
	// files property.
	//
	// Specify an array of string values to match this event if the actual value of files is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
	// totalGb property.
	//
	// Specify an array of string values to match this event if the actual value of totalGb is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TotalGb *[]*string `field:"optional" json:"totalGb" yaml:"totalGb"`
	// volumes property.
	//
	// Specify an array of string values to match this event if the actual value of volumes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

