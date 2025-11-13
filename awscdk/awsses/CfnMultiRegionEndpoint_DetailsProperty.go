package awsses


// An object that contains configuration details of multi-region endpoint (global-endpoint).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   detailsProperty := &DetailsProperty{
//   	RouteDetails: []interface{}{
//   		&RouteDetailsItemsProperty{
//   			Region: jsii.String("region"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-multiregionendpoint-details.html
//
type CfnMultiRegionEndpoint_DetailsProperty struct {
	// A list of route configuration details.
	//
	// Must contain exactly one route configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-multiregionendpoint-details.html#cfn-ses-multiregionendpoint-details-routedetails
	//
	RouteDetails interface{} `field:"required" json:"routeDetails" yaml:"routeDetails"`
}

