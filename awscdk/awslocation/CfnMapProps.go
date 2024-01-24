package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMap`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMapProps := &CfnMapProps{
//   	Configuration: &MapConfigurationProperty{
//   		Style: jsii.String("style"),
//
//   		// the properties below are optional
//   		CustomLayers: []*string{
//   			jsii.String("customLayers"),
//   		},
//   		PoliticalView: jsii.String("politicalView"),
//   	},
//   	MapName: jsii.String("mapName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html
//
type CfnMapProps struct {
	// Specifies the `MapConfiguration` , including the map style, for the map resource that you create.
	//
	// The map style defines the look of maps and the data provider for your map resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html#cfn-location-map-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The name for the map resource.
	//
	// Requirements:
	//
	// - Must contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique map resource name.
	// - No spaces allowed. For example, `ExampleMap` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html#cfn-location-map-mapname
	//
	MapName *string `field:"required" json:"mapName" yaml:"mapName"`
	// An optional description for the map resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html#cfn-location-map-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html#cfn-location-map-pricingplan
	//
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// Applies one or more tags to the map resource.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-map.html#cfn-location-map-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

