package previewawsguarddutyevents


// Type definition for EbsVolumeDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebsVolumeDetails := &EbsVolumeDetails{
//   	ScannedVolumeDetails: []EbsVolumeDetailsItem{
//   		&EbsVolumeDetailsItem{
//   			DeviceName: []*string{
//   				jsii.String("deviceName"),
//   			},
//   			EncryptionType: []*string{
//   				jsii.String("encryptionType"),
//   			},
//   			KmsKeyArn: []*string{
//   				jsii.String("kmsKeyArn"),
//   			},
//   			SnapshotArn: []*string{
//   				jsii.String("snapshotArn"),
//   			},
//   			VolumeArn: []*string{
//   				jsii.String("volumeArn"),
//   			},
//   			VolumeSizeInGb: []*string{
//   				jsii.String("volumeSizeInGb"),
//   			},
//   			VolumeType: []*string{
//   				jsii.String("volumeType"),
//   			},
//   		},
//   	},
//   	SkippedVolumeDetails: []*string{
//   		jsii.String("skippedVolumeDetails"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EbsVolumeDetails struct {
	// scannedVolumeDetails property.
	//
	// Specify an array of string values to match this event if the actual value of scannedVolumeDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScannedVolumeDetails *[]*DetectorEvents_GuardDutyFinding_EbsVolumeDetailsItem `field:"optional" json:"scannedVolumeDetails" yaml:"scannedVolumeDetails"`
	// skippedVolumeDetails property.
	//
	// Specify an array of string values to match this event if the actual value of skippedVolumeDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SkippedVolumeDetails *[]*string `field:"optional" json:"skippedVolumeDetails" yaml:"skippedVolumeDetails"`
}

