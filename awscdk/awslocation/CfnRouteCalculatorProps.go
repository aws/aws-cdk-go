package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouteCalculator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteCalculatorProps := &CfnRouteCalculatorProps{
//   	CalculatorName: jsii.String("calculatorName"),
//   	DataSource: jsii.String("dataSource"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html
//
type CfnRouteCalculatorProps struct {
	// The name of the route calculator resource.
	//
	// Requirements:
	//
	// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique Route calculator resource name.
	// - No spaces allowed. For example, `ExampleRouteCalculator` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html#cfn-location-routecalculator-calculatorname
	//
	CalculatorName *string `field:"required" json:"calculatorName" yaml:"calculatorName"`
	// Specifies the data provider of traffic and road network data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/latest/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://docs.aws.amazon.com/https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm) .
	//
	// Route calculators that use Esri as a data source only calculate routes that are shorter than 400 km.
	// - `Grab` – Grab provides routing functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/latest/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/latest/developerguide/grab.html#grab-coverage-area) .
	// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/latest/developerguide/HERE.html) ' coverage in your region of interest, see [HERE car routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html) .
	//
	// For additional information , see [Data providers](https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html#cfn-location-routecalculator-datasource
	//
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The optional description for the route calculator resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html#cfn-location-routecalculator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html#cfn-location-routecalculator-pricingplan
	//
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-routecalculator.html#cfn-location-routecalculator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

