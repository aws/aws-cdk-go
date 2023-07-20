package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBucketProps := &CfnBucketProps{
//   	BucketName: jsii.String("bucketName"),
//   	BundleId: jsii.String("bundleId"),
//
//   	// the properties below are optional
//   	AccessRules: &AccessRulesProperty{
//   		AllowPublicOverrides: jsii.Boolean(false),
//   		ObjectAccess: jsii.String("objectAccess"),
//   	},
//   	ObjectVersioning: jsii.Boolean(false),
//   	ReadOnlyAccessAccounts: []*string{
//   		jsii.String("readOnlyAccessAccounts"),
//   	},
//   	ResourcesReceivingAccess: []*string{
//   		jsii.String("resourcesReceivingAccess"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html
//
type CfnBucketProps struct {
	// The name of the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bundle ID for the bucket (for example, `small_1_0` ).
	//
	// A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the access rules for the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-accessrules
	//
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// Indicates whether object versioning is enabled for the bucket.
	//
	// The following options can be configured:
	//
	// - `Enabled` - Object versioning is enabled.
	// - `Suspended` - Object versioning was previously enabled but is currently suspended. Existing object versions are retained.
	// - `NeverEnabled` - Object versioning has never been enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-objectversioning
	//
	ObjectVersioning interface{} `field:"optional" json:"objectVersioning" yaml:"objectVersioning"`
	// An array of AWS account IDs that have read-only access to the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-readonlyaccessaccounts
	//
	ReadOnlyAccessAccounts *[]*string `field:"optional" json:"readOnlyAccessAccounts" yaml:"readOnlyAccessAccounts"`
	// An array of Lightsail instances that have access to the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-resourcesreceivingaccess
	//
	ResourcesReceivingAccess *[]*string `field:"optional" json:"resourcesReceivingAccess" yaml:"resourcesReceivingAccess"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html#cfn-lightsail-bucket-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

