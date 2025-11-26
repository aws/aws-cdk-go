package previewawsguarddutyevents


// Type definition for BucketLevelPermissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bucketLevelPermissions := &BucketLevelPermissions{
//   	AccessControlList: &AccessControlList{
//   		AllowsPublicReadAccess: []*string{
//   			jsii.String("allowsPublicReadAccess"),
//   		},
//   		AllowsPublicWriteAccess: []*string{
//   			jsii.String("allowsPublicWriteAccess"),
//   		},
//   	},
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
//   	BucketPolicy: &AccessControlList{
//   		AllowsPublicReadAccess: []*string{
//   			jsii.String("allowsPublicReadAccess"),
//   		},
//   		AllowsPublicWriteAccess: []*string{
//   			jsii.String("allowsPublicWriteAccess"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_BucketLevelPermissions struct {
	// accessControlList property.
	//
	// Specify an array of string values to match this event if the actual value of accessControlList is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccessControlList *DetectorEvents_GuardDutyFinding_AccessControlList `field:"optional" json:"accessControlList" yaml:"accessControlList"`
	// blockPublicAccess property.
	//
	// Specify an array of string values to match this event if the actual value of blockPublicAccess is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockPublicAccess *DetectorEvents_GuardDutyFinding_BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// bucketPolicy property.
	//
	// Specify an array of string values to match this event if the actual value of bucketPolicy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BucketPolicy *DetectorEvents_GuardDutyFinding_AccessControlList `field:"optional" json:"bucketPolicy" yaml:"bucketPolicy"`
}

