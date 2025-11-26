package previewawsguarddutyevents


// Type definition for EbsVolumeDetailsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebsVolumeDetailsItem := &EbsVolumeDetailsItem{
//   	DeviceName: []*string{
//   		jsii.String("deviceName"),
//   	},
//   	EncryptionType: []*string{
//   		jsii.String("encryptionType"),
//   	},
//   	KmsKeyArn: []*string{
//   		jsii.String("kmsKeyArn"),
//   	},
//   	SnapshotArn: []*string{
//   		jsii.String("snapshotArn"),
//   	},
//   	VolumeArn: []*string{
//   		jsii.String("volumeArn"),
//   	},
//   	VolumeSizeInGb: []*string{
//   		jsii.String("volumeSizeInGb"),
//   	},
//   	VolumeType: []*string{
//   		jsii.String("volumeType"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EbsVolumeDetailsItem struct {
	// deviceName property.
	//
	// Specify an array of string values to match this event if the actual value of deviceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeviceName *[]*string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// encryptionType property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionType *[]*string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// kmsKeyArn property.
	//
	// Specify an array of string values to match this event if the actual value of kmsKeyArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KmsKeyArn *[]*string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// snapshotArn property.
	//
	// Specify an array of string values to match this event if the actual value of snapshotArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SnapshotArn *[]*string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// volumeArn property.
	//
	// Specify an array of string values to match this event if the actual value of volumeArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VolumeArn *[]*string `field:"optional" json:"volumeArn" yaml:"volumeArn"`
	// volumeSizeInGB property.
	//
	// Specify an array of string values to match this event if the actual value of volumeSizeInGB is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VolumeSizeInGb *[]*string `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// volumeType property.
	//
	// Specify an array of string values to match this event if the actual value of volumeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VolumeType *[]*string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

