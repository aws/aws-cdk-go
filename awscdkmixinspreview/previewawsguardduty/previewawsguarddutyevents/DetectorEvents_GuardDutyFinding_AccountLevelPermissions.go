package previewawsguarddutyevents


// Type definition for AccountLevelPermissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accountLevelPermissions := &AccountLevelPermissions{
//   	BlockPublicAccess: &BlockPublicAccess{
//   		BlockPublicAcls: []*string{
//   			jsii.String("blockPublicAcls"),
//   		},
//   		BlockPublicPolicy: []*string{
//   			jsii.String("blockPublicPolicy"),
//   		},
//   		IgnorePublicAcls: []*string{
//   			jsii.String("ignorePublicAcls"),
//   		},
//   		RestrictPublicBuckets: []*string{
//   			jsii.String("restrictPublicBuckets"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AccountLevelPermissions struct {
	// blockPublicAccess property.
	//
	// Specify an array of string values to match this event if the actual value of blockPublicAccess is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockPublicAccess *DetectorEvents_GuardDutyFinding_BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
}

