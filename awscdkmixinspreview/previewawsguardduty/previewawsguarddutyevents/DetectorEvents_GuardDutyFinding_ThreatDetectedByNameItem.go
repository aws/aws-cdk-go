package previewawsguarddutyevents


// Type definition for ThreatDetectedByNameItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   threatDetectedByNameItem := &ThreatDetectedByNameItem{
//   	FilePaths: []ThreatDetectedByNameItemItem{
//   		&ThreatDetectedByNameItemItem{
//   			FileName: []*string{
//   				jsii.String("fileName"),
//   			},
//   			FilePath: []*string{
//   				jsii.String("filePath"),
//   			},
//   			Hash: []*string{
//   				jsii.String("hash"),
//   			},
//   			VolumeArn: []*string{
//   				jsii.String("volumeArn"),
//   			},
//   		},
//   	},
//   	ItemCount: []*string{
//   		jsii.String("itemCount"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ThreatDetectedByNameItem struct {
	// filePaths property.
	//
	// Specify an array of string values to match this event if the actual value of filePaths is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePaths *[]*DetectorEvents_GuardDutyFinding_ThreatDetectedByNameItemItem `field:"optional" json:"filePaths" yaml:"filePaths"`
	// itemCount property.
	//
	// Specify an array of string values to match this event if the actual value of itemCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ItemCount *[]*string `field:"optional" json:"itemCount" yaml:"itemCount"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
}

