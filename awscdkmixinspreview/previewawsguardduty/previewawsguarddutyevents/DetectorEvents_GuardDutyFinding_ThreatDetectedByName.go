package previewawsguarddutyevents


// Type definition for ThreatDetectedByName.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   threatDetectedByName := &ThreatDetectedByName{
//   	ItemCount: []*string{
//   		jsii.String("itemCount"),
//   	},
//   	Shortened: []*string{
//   		jsii.String("shortened"),
//   	},
//   	ThreatNames: []ThreatDetectedByNameItem{
//   		&ThreatDetectedByNameItem{
//   			FilePaths: []ThreatDetectedByNameItemItem{
//   				&ThreatDetectedByNameItemItem{
//   					FileName: []*string{
//   						jsii.String("fileName"),
//   					},
//   					FilePath: []*string{
//   						jsii.String("filePath"),
//   					},
//   					Hash: []*string{
//   						jsii.String("hash"),
//   					},
//   					VolumeArn: []*string{
//   						jsii.String("volumeArn"),
//   					},
//   				},
//   			},
//   			ItemCount: []*string{
//   				jsii.String("itemCount"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Severity: []*string{
//   				jsii.String("severity"),
//   			},
//   		},
//   	},
//   	UniqueThreatNameCount: []*string{
//   		jsii.String("uniqueThreatNameCount"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ThreatDetectedByName struct {
	// itemCount property.
	//
	// Specify an array of string values to match this event if the actual value of itemCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ItemCount *[]*string `field:"optional" json:"itemCount" yaml:"itemCount"`
	// shortened property.
	//
	// Specify an array of string values to match this event if the actual value of shortened is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Shortened *[]*string `field:"optional" json:"shortened" yaml:"shortened"`
	// threatNames property.
	//
	// Specify an array of string values to match this event if the actual value of threatNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatNames *[]*DetectorEvents_GuardDutyFinding_ThreatDetectedByNameItem `field:"optional" json:"threatNames" yaml:"threatNames"`
	// uniqueThreatNameCount property.
	//
	// Specify an array of string values to match this event if the actual value of uniqueThreatNameCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UniqueThreatNameCount *[]*string `field:"optional" json:"uniqueThreatNameCount" yaml:"uniqueThreatNameCount"`
}

