package previewawsguarddutyevents


// Type definition for ResourceItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceItem := &ResourceItem{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	DefaultServerSideEncryption: &DefaultServerSideEncryption{
//   		EncryptionType: []*string{
//   			jsii.String("encryptionType"),
//   		},
//   		KmsMasterKeyArn: []*string{
//   			jsii.String("kmsMasterKeyArn"),
//   		},
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Owner: &Owner{
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   	},
//   	PublicAccess: &PublicAccess{
//   		EffectivePermission: []*string{
//   			jsii.String("effectivePermission"),
//   		},
//   		PermissionConfiguration: &PermissionConfiguration{
//   			AccountLevelPermissions: &AccountLevelPermissions{
//   				BlockPublicAccess: &BlockPublicAccess{
//   					BlockPublicAcls: []*string{
//   						jsii.String("blockPublicAcls"),
//   					},
//   					BlockPublicPolicy: []*string{
//   						jsii.String("blockPublicPolicy"),
//   					},
//   					IgnorePublicAcls: []*string{
//   						jsii.String("ignorePublicAcls"),
//   					},
//   					RestrictPublicBuckets: []*string{
//   						jsii.String("restrictPublicBuckets"),
//   					},
//   				},
//   			},
//   			BucketLevelPermissions: &BucketLevelPermissions{
//   				AccessControlList: &AccessControlList{
//   					AllowsPublicReadAccess: []*string{
//   						jsii.String("allowsPublicReadAccess"),
//   					},
//   					AllowsPublicWriteAccess: []*string{
//   						jsii.String("allowsPublicWriteAccess"),
//   					},
//   				},
//   				BlockPublicAccess: &BlockPublicAccess{
//   					BlockPublicAcls: []*string{
//   						jsii.String("blockPublicAcls"),
//   					},
//   					BlockPublicPolicy: []*string{
//   						jsii.String("blockPublicPolicy"),
//   					},
//   					IgnorePublicAcls: []*string{
//   						jsii.String("ignorePublicAcls"),
//   					},
//   					RestrictPublicBuckets: []*string{
//   						jsii.String("restrictPublicBuckets"),
//   					},
//   				},
//   				BucketPolicy: &AccessControlList{
//   					AllowsPublicReadAccess: []*string{
//   						jsii.String("allowsPublicReadAccess"),
//   					},
//   					AllowsPublicWriteAccess: []*string{
//   						jsii.String("allowsPublicWriteAccess"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []EcsClusterDetailsItem{
//   		&EcsClusterDetailsItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ResourceItem struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// defaultServerSideEncryption property.
	//
	// Specify an array of string values to match this event if the actual value of defaultServerSideEncryption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefaultServerSideEncryption *DetectorEvents_GuardDutyFinding_DefaultServerSideEncryption `field:"optional" json:"defaultServerSideEncryption" yaml:"defaultServerSideEncryption"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// owner property.
	//
	// Specify an array of string values to match this event if the actual value of owner is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Owner *DetectorEvents_GuardDutyFinding_Owner `field:"optional" json:"owner" yaml:"owner"`
	// publicAccess property.
	//
	// Specify an array of string values to match this event if the actual value of publicAccess is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PublicAccess *DetectorEvents_GuardDutyFinding_PublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*DetectorEvents_GuardDutyFinding_EcsClusterDetailsItem `field:"optional" json:"tags" yaml:"tags"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

