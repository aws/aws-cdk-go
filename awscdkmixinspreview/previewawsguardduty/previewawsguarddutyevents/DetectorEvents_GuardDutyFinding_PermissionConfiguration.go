package previewawsguarddutyevents


// Type definition for PermissionConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   permissionConfiguration := &PermissionConfiguration{
//   	AccountLevelPermissions: &AccountLevelPermissions{
//   		BlockPublicAccess: &BlockPublicAccess{
//   			BlockPublicAcls: []*string{
//   				jsii.String("blockPublicAcls"),
//   			},
//   			BlockPublicPolicy: []*string{
//   				jsii.String("blockPublicPolicy"),
//   			},
//   			IgnorePublicAcls: []*string{
//   				jsii.String("ignorePublicAcls"),
//   			},
//   			RestrictPublicBuckets: []*string{
//   				jsii.String("restrictPublicBuckets"),
//   			},
//   		},
//   	},
//   	BucketLevelPermissions: &BucketLevelPermissions{
//   		AccessControlList: &AccessControlList{
//   			AllowsPublicReadAccess: []*string{
//   				jsii.String("allowsPublicReadAccess"),
//   			},
//   			AllowsPublicWriteAccess: []*string{
//   				jsii.String("allowsPublicWriteAccess"),
//   			},
//   		},
//   		BlockPublicAccess: &BlockPublicAccess{
//   			BlockPublicAcls: []*string{
//   				jsii.String("blockPublicAcls"),
//   			},
//   			BlockPublicPolicy: []*string{
//   				jsii.String("blockPublicPolicy"),
//   			},
//   			IgnorePublicAcls: []*string{
//   				jsii.String("ignorePublicAcls"),
//   			},
//   			RestrictPublicBuckets: []*string{
//   				jsii.String("restrictPublicBuckets"),
//   			},
//   		},
//   		BucketPolicy: &AccessControlList{
//   			AllowsPublicReadAccess: []*string{
//   				jsii.String("allowsPublicReadAccess"),
//   			},
//   			AllowsPublicWriteAccess: []*string{
//   				jsii.String("allowsPublicWriteAccess"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_PermissionConfiguration struct {
	// accountLevelPermissions property.
	//
	// Specify an array of string values to match this event if the actual value of accountLevelPermissions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountLevelPermissions *DetectorEvents_GuardDutyFinding_AccountLevelPermissions `field:"optional" json:"accountLevelPermissions" yaml:"accountLevelPermissions"`
	// bucketLevelPermissions property.
	//
	// Specify an array of string values to match this event if the actual value of bucketLevelPermissions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BucketLevelPermissions *DetectorEvents_GuardDutyFinding_BucketLevelPermissions `field:"optional" json:"bucketLevelPermissions" yaml:"bucketLevelPermissions"`
}

