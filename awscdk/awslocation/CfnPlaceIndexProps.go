package awslocation


// Properties for defining a `CfnPlaceIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlaceIndexProps := &CfnPlaceIndexProps{
//   	DataSource: jsii.String("dataSource"),
//   	IndexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		IntendedUse: jsii.String("intendedUse"),
//   	},
//   	Description: jsii.String("description"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html
//
type CfnPlaceIndexProps struct {
	// Specifies the geospatial data provider for the new place index.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/latest/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on geocoding coverage](https://docs.aws.amazon.com/https://developers.arcgis.com/rest/geocode/api-reference/geocode-coverage.htm) .
	// - `Grab` – Grab provides place index functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/latest/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/latest/developerguide/grab.html#grab-coverage-area) .
	// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/latest/developerguide/HERE.html) ' coverage in your region of interest, see [HERE details on goecoding coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/geocoder/dev_guide/topics/coverage-geocoder.html) .
	//
	// > If you specify HERE Technologies ( `Here` ) as the data provider, you may not [store results](https://docs.aws.amazon.com//location-places/latest/APIReference/API_DataSourceConfiguration.html) for locations in Japan. For more information, see the [AWS Service Terms](https://docs.aws.amazon.com/service-terms/) for Amazon Location Service.
	//
	// For additional information , see [Data providers](https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html#cfn-location-placeindex-datasource
	//
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of the place index resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique place index resource name.
	// - No spaces allowed. For example, `ExamplePlaceIndex` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html#cfn-location-placeindex-indexname
	//
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Specifies the data storage option requesting Places.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html#cfn-location-placeindex-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The optional description for the place index resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html#cfn-location-placeindex-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-placeindex.html#cfn-location-placeindex-pricingplan
	//
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
}

