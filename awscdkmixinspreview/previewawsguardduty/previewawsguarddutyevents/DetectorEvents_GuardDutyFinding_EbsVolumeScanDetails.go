package previewawsguarddutyevents


// Type definition for EbsVolumeScanDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebsVolumeScanDetails := &EbsVolumeScanDetails{
//   	ScanCompletedAt: []*string{
//   		jsii.String("scanCompletedAt"),
//   	},
//   	ScanDetections: &ScanDetections{
//   		HighestSeverityThreatDetails: &HighestSeverityThreatDetails{
//   			Count: []*string{
//   				jsii.String("count"),
//   			},
//   			Severity: []*string{
//   				jsii.String("severity"),
//   			},
//   			ThreatName: []*string{
//   				jsii.String("threatName"),
//   			},
//   		},
//   		ScannedItemCount: &ScannedItemCount{
//   			Files: []*string{
//   				jsii.String("files"),
//   			},
//   			TotalGb: []*string{
//   				jsii.String("totalGb"),
//   			},
//   			Volumes: []*string{
//   				jsii.String("volumes"),
//   			},
//   		},
//   		ThreatDetectedByName: &ThreatDetectedByName{
//   			ItemCount: []*string{
//   				jsii.String("itemCount"),
//   			},
//   			Shortened: []*string{
//   				jsii.String("shortened"),
//   			},
//   			ThreatNames: []ThreatDetectedByNameItem{
//   				&ThreatDetectedByNameItem{
//   					FilePaths: []ThreatDetectedByNameItemItem{
//   						&ThreatDetectedByNameItemItem{
//   							FileName: []*string{
//   								jsii.String("fileName"),
//   							},
//   							FilePath: []*string{
//   								jsii.String("filePath"),
//   							},
//   							Hash: []*string{
//   								jsii.String("hash"),
//   							},
//   							VolumeArn: []*string{
//   								jsii.String("volumeArn"),
//   							},
//   						},
//   					},
//   					ItemCount: []*string{
//   						jsii.String("itemCount"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Severity: []*string{
//   						jsii.String("severity"),
//   					},
//   				},
//   			},
//   			UniqueThreatNameCount: []*string{
//   				jsii.String("uniqueThreatNameCount"),
//   			},
//   		},
//   		ThreatsDetectedItemCount: &ThreatsDetectedItemCount{
//   			Files: []*string{
//   				jsii.String("files"),
//   			},
//   		},
//   	},
//   	ScanId: []*string{
//   		jsii.String("scanId"),
//   	},
//   	ScanStartedAt: []*string{
//   		jsii.String("scanStartedAt"),
//   	},
//   	Sources: []*string{
//   		jsii.String("sources"),
//   	},
//   	TriggerFindingId: []*string{
//   		jsii.String("triggerFindingId"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EbsVolumeScanDetails struct {
	// scanCompletedAt property.
	//
	// Specify an array of string values to match this event if the actual value of scanCompletedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScanCompletedAt *[]*string `field:"optional" json:"scanCompletedAt" yaml:"scanCompletedAt"`
	// scanDetections property.
	//
	// Specify an array of string values to match this event if the actual value of scanDetections is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScanDetections *DetectorEvents_GuardDutyFinding_ScanDetections `field:"optional" json:"scanDetections" yaml:"scanDetections"`
	// scanId property.
	//
	// Specify an array of string values to match this event if the actual value of scanId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScanId *[]*string `field:"optional" json:"scanId" yaml:"scanId"`
	// scanStartedAt property.
	//
	// Specify an array of string values to match this event if the actual value of scanStartedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScanStartedAt *[]*string `field:"optional" json:"scanStartedAt" yaml:"scanStartedAt"`
	// sources property.
	//
	// Specify an array of string values to match this event if the actual value of sources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
	// triggerFindingId property.
	//
	// Specify an array of string values to match this event if the actual value of triggerFindingId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TriggerFindingId *[]*string `field:"optional" json:"triggerFindingId" yaml:"triggerFindingId"`
}

