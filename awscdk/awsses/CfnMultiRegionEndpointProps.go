package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMultiRegionEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiRegionEndpointProps := &CfnMultiRegionEndpointProps{
//   	Details: &DetailsProperty{
//   		RouteDetails: []interface{}{
//   			&RouteDetailsItemsProperty{
//   				Region: jsii.String("region"),
//   			},
//   		},
//   	},
//   	EndpointName: jsii.String("endpointName"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-multiregionendpoint.html
//
type CfnMultiRegionEndpointProps struct {
	// Contains details of a multi-region endpoint (global-endpoint) being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-multiregionendpoint.html#cfn-ses-multiregionendpoint-details
	//
	Details interface{} `field:"required" json:"details" yaml:"details"`
	// The name of the multi-region endpoint (global-endpoint).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-multiregionendpoint.html#cfn-ses-multiregionendpoint-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// An array of objects that define the tags (keys and values) to associate with the multi-region endpoint (global-endpoint).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-multiregionendpoint.html#cfn-ses-multiregionendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

