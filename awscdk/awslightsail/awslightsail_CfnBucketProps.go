package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBucketProps := &cfnBucketProps{
//   	bucketName: jsii.String("bucketName"),
//   	bundleId: jsii.String("bundleId"),
//
//   	// the properties below are optional
//   	accessRules: &accessRulesProperty{
//   		allowPublicOverrides: jsii.Boolean(false),
//   		objectAccess: jsii.String("objectAccess"),
//   	},
//   	objectVersioning: jsii.Boolean(false),
//   	readOnlyAccessAccounts: []*string{
//   		jsii.String("readOnlyAccessAccounts"),
//   	},
//   	resourcesReceivingAccess: []*string{
//   		jsii.String("resourcesReceivingAccess"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucketProps struct {
	// The name of the bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bundle ID for the bucket (for example, `small_1_0` ).
	//
	// A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a bucket.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the access rules for the bucket.
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// Indicates whether object versioning is enabled for the bucket.
	//
	// The following options can be configured:
	//
	// - `Enabled` - Object versioning is enabled.
	// - `Suspended` - Object versioning was previously enabled but is currently suspended. Existing object versions are retained.
	// - `NeverEnabled` - Object versioning has never been enabled.
	ObjectVersioning interface{} `field:"optional" json:"objectVersioning" yaml:"objectVersioning"`
	// An array of AWS account IDs that have read-only access to the bucket.
	ReadOnlyAccessAccounts *[]*string `field:"optional" json:"readOnlyAccessAccounts" yaml:"readOnlyAccessAccounts"`
	// An array of Lightsail instances that have access to the bucket.
	ResourcesReceivingAccess *[]*string `field:"optional" json:"resourcesReceivingAccess" yaml:"resourcesReceivingAccess"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

