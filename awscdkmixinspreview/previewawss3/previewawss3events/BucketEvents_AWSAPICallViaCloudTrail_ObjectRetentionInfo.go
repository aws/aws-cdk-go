package previewawss3events


// Type definition for ObjectRetentionInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRetentionInfo := &ObjectRetentionInfo{
//   	LegalHoldInfo: &LegalHoldInfo{
//   		IsUnderLegalHold: []*string{
//   			jsii.String("isUnderLegalHold"),
//   		},
//   		LastModifiedTime: []*string{
//   			jsii.String("lastModifiedTime"),
//   		},
//   	},
//   	RetentionInfo: &RetentionInfo{
//   		LastModifiedTime: []*string{
//   			jsii.String("lastModifiedTime"),
//   		},
//   		RetainUntilMode: []*string{
//   			jsii.String("retainUntilMode"),
//   		},
//   		RetainUntilTime: []*string{
//   			jsii.String("retainUntilTime"),
//   		},
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_ObjectRetentionInfo struct {
	// legalHoldInfo property.
	//
	// Specify an array of string values to match this event if the actual value of legalHoldInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LegalHoldInfo *BucketEvents_AWSAPICallViaCloudTrail_LegalHoldInfo `field:"optional" json:"legalHoldInfo" yaml:"legalHoldInfo"`
	// retentionInfo property.
	//
	// Specify an array of string values to match this event if the actual value of retentionInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetentionInfo *BucketEvents_AWSAPICallViaCloudTrail_RetentionInfo `field:"optional" json:"retentionInfo" yaml:"retentionInfo"`
}

