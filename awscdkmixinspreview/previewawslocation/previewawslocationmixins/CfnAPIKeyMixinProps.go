package previewawslocationmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAPIKeyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAPIKeyMixinProps := &CfnAPIKeyMixinProps{
//   	Description: jsii.String("description"),
//   	ExpireTime: jsii.String("expireTime"),
//   	ForceDelete: jsii.Boolean(false),
//   	ForceUpdate: jsii.Boolean(false),
//   	KeyName: jsii.String("keyName"),
//   	NoExpiry: jsii.Boolean(false),
//   	Restrictions: &ApiKeyRestrictionsProperty{
//   		AllowActions: []*string{
//   			jsii.String("allowActions"),
//   		},
//   		AllowAndroidApps: []interface{}{
//   			&AndroidAppProperty{
//   				CertificateFingerprint: jsii.String("certificateFingerprint"),
//   				Package: jsii.String("package"),
//   			},
//   		},
//   		AllowAppleApps: []interface{}{
//   			&AppleAppProperty{
//   				BundleId: jsii.String("bundleId"),
//   			},
//   		},
//   		AllowReferers: []*string{
//   			jsii.String("allowReferers"),
//   		},
//   		AllowResources: []*string{
//   			jsii.String("allowResources"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html
//
type CfnAPIKeyMixinProps struct {
	// Updates the description for the API key resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The optional timestamp for when the API key resource will expire in [ISO 8601 format](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-expiretime
	//
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// ForceDelete bypasses an API key's expiry conditions and deletes the key.
	//
	// Set the parameter `true` to delete the key or to `false` to not preemptively delete the API key.
	//
	// Valid values: `true` , or `false` .
	//
	// > This action is irreversible. Only use ForceDelete if you are certain the key is no longer in use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-forcedelete
	//
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// The boolean flag to be included for updating `ExpireTime` or Restrictions details.
	//
	// Must be set to `true` to update an API key resource that has been used in the past 7 days. `False` if force update is not preferred.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-forceupdate
	//
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// A custom name for the API key resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique API key name.
	// - No spaces allowed. For example, `ExampleAPIKey` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-keyname
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Whether the API key should expire.
	//
	// Set to `true` to set the API key to have no expiration time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-noexpiry
	//
	NoExpiry interface{} `field:"optional" json:"noExpiry" yaml:"noExpiry"`
	// The API key restrictions for the API key resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-restrictions
	//
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Applies one or more tags to the map resource.
	//
	// A tag is a key-value pair that helps manage, identify, search, and filter your resources by labelling them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html#cfn-location-apikey-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

