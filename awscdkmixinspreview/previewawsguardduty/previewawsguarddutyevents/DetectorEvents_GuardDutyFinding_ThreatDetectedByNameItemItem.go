package previewawsguarddutyevents


// Type definition for ThreatDetectedByNameItemItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   threatDetectedByNameItemItem := &ThreatDetectedByNameItemItem{
//   	FileName: []*string{
//   		jsii.String("fileName"),
//   	},
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	Hash: []*string{
//   		jsii.String("hash"),
//   	},
//   	VolumeArn: []*string{
//   		jsii.String("volumeArn"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ThreatDetectedByNameItemItem struct {
	// fileName property.
	//
	// Specify an array of string values to match this event if the actual value of fileName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileName *[]*string `field:"optional" json:"fileName" yaml:"fileName"`
	// filePath property.
	//
	// Specify an array of string values to match this event if the actual value of filePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// hash property.
	//
	// Specify an array of string values to match this event if the actual value of hash is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Hash *[]*string `field:"optional" json:"hash" yaml:"hash"`
	// volumeArn property.
	//
	// Specify an array of string values to match this event if the actual value of volumeArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VolumeArn *[]*string `field:"optional" json:"volumeArn" yaml:"volumeArn"`
}

