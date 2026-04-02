package awsmediapackagev2alpha


// The base URLs to use for retrieving segments.
//
// You can specify multiple locations and indicate the
// priority and weight for when each should be used, for use in multi-CDN workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashBaseUrlProperty := &DashBaseUrlProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	DvbPriority: jsii.Number(123),
//   	DvbWeight: jsii.Number(123),
//   	ServiceLocation: jsii.String("serviceLocation"),
//   }
//
// Experimental.
type DashBaseUrlProperty struct {
	// A source location for segments.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// For use with DVB-DASH profiles only.
	//
	// The priority of this location for serving segments. The lower the number, the higher the priority.
	// Default: - No priority specified.
	//
	// Experimental.
	DvbPriority *float64 `field:"optional" json:"dvbPriority" yaml:"dvbPriority"`
	// For use with DVB-DASH profiles only.
	//
	// The weighting for source locations that have the same priority.
	// Default: - No weight specified.
	//
	// Experimental.
	DvbWeight *float64 `field:"optional" json:"dvbWeight" yaml:"dvbWeight"`
	// The name of the source location.
	// Default: - No service location specified.
	//
	// Experimental.
	ServiceLocation *string `field:"optional" json:"serviceLocation" yaml:"serviceLocation"`
}

