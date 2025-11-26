package previewawslocationmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTrackerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrackerMixinProps := &CfnTrackerMixinProps{
//   	Description: jsii.String("description"),
//   	EventBridgeEnabled: jsii.Boolean(false),
//   	KmsKeyEnableGeospatialQueries: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PositionFiltering: jsii.String("positionFiltering"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   	PricingPlanDataSource: jsii.String("pricingPlanDataSource"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackerName: jsii.String("trackerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html
//
type CfnTrackerMixinProps struct {
	// An optional description for the tracker resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-eventbridgeenabled
	//
	EventBridgeEnabled interface{} `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-kmskeyenablegeospatialqueries
	//
	KmsKeyEnableGeospatialQueries interface{} `field:"optional" json:"kmsKeyEnableGeospatialQueries" yaml:"kmsKeyEnableGeospatialQueries"`
	// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the position filtering for the tracker resource.
	//
	// Valid values:
	//
	// - `TimeBased` - Location updates are evaluated against linked geofence collections, but not every location update is stored. If your update frequency is more often than 30 seconds, only one update per 30 seconds is stored for each unique device ID.
	// - `DistanceBased` - If the device has moved less than 30 m (98.4 ft), location updates are ignored. Location updates within this area are neither evaluated against linked geofence collections, nor stored. This helps control costs by reducing the number of geofence evaluations and historical device positions to paginate through. Distance-based filtering can also reduce the effects of GPS noise when displaying device trajectories on a map.
	// - `AccuracyBased` - If the device has moved less than the measured accuracy, location updates are ignored. For example, if two consecutive updates from a device have a horizontal accuracy of 5 m and 10 m, the second update is ignored if the device has moved less than 15 m. Ignored location updates are neither evaluated against linked geofence collections, nor stored. This can reduce the effects of GPS noise when displaying device trajectories on a map, and can help control your costs by reducing the number of geofence evaluations.
	//
	// This field is optional. If not specified, the default value is `TimeBased` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-positionfiltering
	//
	PositionFiltering *string `field:"optional" json:"positionFiltering" yaml:"positionFiltering"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-pricingplan
	//
	// Deprecated: this property has been deprecated.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// This shape is deprecated since 2022-02-01: Deprecated.
	//
	// No longer allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-pricingplandatasource
	//
	// Deprecated: this property has been deprecated.
	PricingPlanDataSource *string `field:"optional" json:"pricingPlanDataSource" yaml:"pricingPlanDataSource"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name for the tracker resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique tracker resource name.
	// - No spaces allowed. For example, `ExampleTracker` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-tracker.html#cfn-location-tracker-trackername
	//
	TrackerName *string `field:"optional" json:"trackerName" yaml:"trackerName"`
}

