package previewawsguarddutyevents


// Type definition for Evidence.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evidence := &Evidence{
//   	ThreatIntelligenceDetails: []EvidenceItem{
//   		&EvidenceItem{
//   			ThreatListName: []*string{
//   				jsii.String("threatListName"),
//   			},
//   			ThreatNames: []*string{
//   				jsii.String("threatNames"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Evidence struct {
	// threatIntelligenceDetails property.
	//
	// Specify an array of string values to match this event if the actual value of threatIntelligenceDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatIntelligenceDetails *[]*DetectorEvents_GuardDutyFinding_EvidenceItem `field:"optional" json:"threatIntelligenceDetails" yaml:"threatIntelligenceDetails"`
}

