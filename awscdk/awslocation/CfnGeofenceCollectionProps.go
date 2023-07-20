package awslocation


// Properties for defining a `CfnGeofenceCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGeofenceCollectionProps := &CfnGeofenceCollectionProps{
//   	CollectionName: jsii.String("collectionName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   	PricingPlanDataSource: jsii.String("pricingPlanDataSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html
//
type CfnGeofenceCollectionProps struct {
	// A custom name for the geofence collection.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique geofence collection name.
	// - No spaces allowed. For example, `ExampleGeofenceCollection` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-collectionname
	//
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-geofencecollection.html#cfn-location-geofencecollection-pricingplandatasource
	//
	// Deprecated: this property has been deprecated.
	PricingPlanDataSource *string `field:"optional" json:"pricingPlanDataSource" yaml:"pricingPlanDataSource"`
}

