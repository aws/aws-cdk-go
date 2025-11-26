package previewawsguarddutyevents


// Type definition for ScanDetections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scanDetections := &ScanDetections{
//   	HighestSeverityThreatDetails: &HighestSeverityThreatDetails{
//   		Count: []*string{
//   			jsii.String("count"),
//   		},
//   		Severity: []*string{
//   			jsii.String("severity"),
//   		},
//   		ThreatName: []*string{
//   			jsii.String("threatName"),
//   		},
//   	},
//   	ScannedItemCount: &ScannedItemCount{
//   		Files: []*string{
//   			jsii.String("files"),
//   		},
//   		TotalGb: []*string{
//   			jsii.String("totalGb"),
//   		},
//   		Volumes: []*string{
//   			jsii.String("volumes"),
//   		},
//   	},
//   	ThreatDetectedByName: &ThreatDetectedByName{
//   		ItemCount: []*string{
//   			jsii.String("itemCount"),
//   		},
//   		Shortened: []*string{
//   			jsii.String("shortened"),
//   		},
//   		ThreatNames: []ThreatDetectedByNameItem{
//   			&ThreatDetectedByNameItem{
//   				FilePaths: []ThreatDetectedByNameItemItem{
//   					&ThreatDetectedByNameItemItem{
//   						FileName: []*string{
//   							jsii.String("fileName"),
//   						},
//   						FilePath: []*string{
//   							jsii.String("filePath"),
//   						},
//   						Hash: []*string{
//   							jsii.String("hash"),
//   						},
//   						VolumeArn: []*string{
//   							jsii.String("volumeArn"),
//   						},
//   					},
//   				},
//   				ItemCount: []*string{
//   					jsii.String("itemCount"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Severity: []*string{
//   					jsii.String("severity"),
//   				},
//   			},
//   		},
//   		UniqueThreatNameCount: []*string{
//   			jsii.String("uniqueThreatNameCount"),
//   		},
//   	},
//   	ThreatsDetectedItemCount: &ThreatsDetectedItemCount{
//   		Files: []*string{
//   			jsii.String("files"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ScanDetections struct {
	// highestSeverityThreatDetails property.
	//
	// Specify an array of string values to match this event if the actual value of highestSeverityThreatDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HighestSeverityThreatDetails *DetectorEvents_GuardDutyFinding_HighestSeverityThreatDetails `field:"optional" json:"highestSeverityThreatDetails" yaml:"highestSeverityThreatDetails"`
	// scannedItemCount property.
	//
	// Specify an array of string values to match this event if the actual value of scannedItemCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScannedItemCount *DetectorEvents_GuardDutyFinding_ScannedItemCount `field:"optional" json:"scannedItemCount" yaml:"scannedItemCount"`
	// threatDetectedByName property.
	//
	// Specify an array of string values to match this event if the actual value of threatDetectedByName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatDetectedByName *DetectorEvents_GuardDutyFinding_ThreatDetectedByName `field:"optional" json:"threatDetectedByName" yaml:"threatDetectedByName"`
	// threatsDetectedItemCount property.
	//
	// Specify an array of string values to match this event if the actual value of threatsDetectedItemCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ThreatsDetectedItemCount *DetectorEvents_GuardDutyFinding_ThreatsDetectedItemCount `field:"optional" json:"threatsDetectedItemCount" yaml:"threatsDetectedItemCount"`
}

