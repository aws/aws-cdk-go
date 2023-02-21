package awslocation


// Properties for defining a `CfnTracker`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrackerProps := &cfnTrackerProps{
//   	trackerName: jsii.String("trackerName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	positionFiltering: jsii.String("positionFiltering"),
//   	pricingPlan: jsii.String("pricingPlan"),
//   	pricingPlanDataSource: jsii.String("pricingPlanDataSource"),
//   }
//
type CfnTrackerProps struct {
	// The name for the tracker resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique tracker resource name.
	// - No spaces allowed. For example, `ExampleTracker` .
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// An optional description for the tracker resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
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
	PositionFiltering *string `field:"optional" json:"positionFiltering" yaml:"positionFiltering"`
	// No longer used.
	//
	// If included, the only allowed value is `RequestBasedUsage` .
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// This parameter is no longer used.
	PricingPlanDataSource *string `field:"optional" json:"pricingPlanDataSource" yaml:"pricingPlanDataSource"`
}

