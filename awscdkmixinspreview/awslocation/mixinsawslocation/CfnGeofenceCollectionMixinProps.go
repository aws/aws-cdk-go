package mixinsawslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnGeofenceCollectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGeofenceCollectionMixinProps := &CfnGeofenceCollectionMixinProps{
//   	CollectionName: jsii.String("collectionName"),
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   	PricingPlanDataSource: jsii.String("pricingPlanDataSource"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html
//
type CfnGeofenceCollectionMixinProps struct {
	// A custom name for the geofence collection.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique geofence collection name.
	// - No spaces allowed. For example, `ExampleGeofenceCollection` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
	//
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// An optional description for the geofence collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplan
	//
	// Deprecated: this property has been deprecated.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// This shape is deprecated since 2022-02-01: Deprecated.
	//
	// No longer allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
	//
	// Deprecated: this property has been deprecated.
	PricingPlanDataSource *string `field:"optional" json:"pricingPlanDataSource" yaml:"pricingPlanDataSource"`
	// Applies one or more tags to the geofence collection.
	//
	// A tag is a key-value pair helps manage, identify, search, and filter your resources by labelling them.
	//
	// Format: `"key" : "value"`
	//
	// Restrictions:
	//
	// - Maximum 50 tags per resource
	// - Each resource tag must be unique with a maximum of one value.
	// - Maximum key length: 128 Unicode characters in UTF-8
	// - Maximum value length: 256 Unicode characters in UTF-8
	// - Can use alphanumeric characters (A–Z, a–z, 0–9), and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

