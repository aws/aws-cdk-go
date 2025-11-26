package previewawsguarddutyevents


// Type definition for EvidenceItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evidenceItem := &EvidenceItem{
//   	ThreatListName: []*string{
//   		jsii.String("threatListName"),
//   	},
//   	ThreatNames: []*string{
//   		jsii.String("threatNames"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EvidenceItem struct {
	// threatListName property.
	//
	// Specify an array of string values to match this event if the actual value of threatListName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatListName *[]*string `field:"optional" json:"threatListName" yaml:"threatListName"`
	// threatNames property.
	//
	// Specify an array of string values to match this event if the actual value of threatNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatNames *[]*string `field:"optional" json:"threatNames" yaml:"threatNames"`
}

