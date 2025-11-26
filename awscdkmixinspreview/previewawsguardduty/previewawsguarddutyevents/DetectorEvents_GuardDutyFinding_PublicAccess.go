package previewawsguarddutyevents


// Type definition for PublicAccess.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicAccess := &PublicAccess{
//   	EffectivePermission: []*string{
//   		jsii.String("effectivePermission"),
//   	},
//   	PermissionConfiguration: &PermissionConfiguration{
//   		AccountLevelPermissions: &AccountLevelPermissions{
//   			BlockPublicAccess: &BlockPublicAccess{
//   				BlockPublicAcls: []*string{
//   					jsii.String("blockPublicAcls"),
//   				},
//   				BlockPublicPolicy: []*string{
//   					jsii.String("blockPublicPolicy"),
//   				},
//   				IgnorePublicAcls: []*string{
//   					jsii.String("ignorePublicAcls"),
//   				},
//   				RestrictPublicBuckets: []*string{
//   					jsii.String("restrictPublicBuckets"),
//   				},
//   			},
//   		},
//   		BucketLevelPermissions: &BucketLevelPermissions{
//   			AccessControlList: &AccessControlList{
//   				AllowsPublicReadAccess: []*string{
//   					jsii.String("allowsPublicReadAccess"),
//   				},
//   				AllowsPublicWriteAccess: []*string{
//   					jsii.String("allowsPublicWriteAccess"),
//   				},
//   			},
//   			BlockPublicAccess: &BlockPublicAccess{
//   				BlockPublicAcls: []*string{
//   					jsii.String("blockPublicAcls"),
//   				},
//   				BlockPublicPolicy: []*string{
//   					jsii.String("blockPublicPolicy"),
//   				},
//   				IgnorePublicAcls: []*string{
//   					jsii.String("ignorePublicAcls"),
//   				},
//   				RestrictPublicBuckets: []*string{
//   					jsii.String("restrictPublicBuckets"),
//   				},
//   			},
//   			BucketPolicy: &AccessControlList{
//   				AllowsPublicReadAccess: []*string{
//   					jsii.String("allowsPublicReadAccess"),
//   				},
//   				AllowsPublicWriteAccess: []*string{
//   					jsii.String("allowsPublicWriteAccess"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_PublicAccess struct {
	// effectivePermission property.
	//
	// Specify an array of string values to match this event if the actual value of effectivePermission is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EffectivePermission *[]*string `field:"optional" json:"effectivePermission" yaml:"effectivePermission"`
	// permissionConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of permissionConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PermissionConfiguration *DetectorEvents_GuardDutyFinding_PermissionConfiguration `field:"optional" json:"permissionConfiguration" yaml:"permissionConfiguration"`
}

