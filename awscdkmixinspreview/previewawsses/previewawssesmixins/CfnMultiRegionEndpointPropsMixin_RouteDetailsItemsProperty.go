package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routeDetailsItemsProperty := &RouteDetailsItemsProperty{
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-multiregionendpoint-routedetailsitems.html
//
type CfnMultiRegionEndpointPropsMixin_RouteDetailsItemsProperty struct {
	// The name of an AWS-Region to be a secondary region for the multi-region endpoint (global-endpoint).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-multiregionendpoint-routedetailsitems.html#cfn-ses-multiregionendpoint-routedetailsitems-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

