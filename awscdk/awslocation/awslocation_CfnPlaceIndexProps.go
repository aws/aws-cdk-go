package awslocation


// Properties for defining a `CfnPlaceIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlaceIndexProps := &cfnPlaceIndexProps{
//   	dataSource: jsii.String("dataSource"),
//   	indexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	dataSourceConfiguration: &dataSourceConfigurationProperty{
//   		intendedUse: jsii.String("intendedUse"),
//   	},
//   	description: jsii.String("description"),
//   	pricingPlan: jsii.String("pricingPlan"),
//   }
//
type CfnPlaceIndexProps struct {
	// Specifies the data provider of geospatial data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` will return an error.
	//
	// Valid values include:
	//
	// - `Esri`
	// - `Here`
	//
	// > Place index resources using HERE as a data provider can't be used to [store](https://docs.aws.amazon.com/location-places/latest/APIReference/API_DataSourceConfiguration.html) results for locations in Japan. For more information, see the [AWS Service Terms](https://docs.aws.amazon.com/service-terms/) for Amazon Location Service.
	//
	// For additional details on data providers, see the [Amazon Location Service data providers page](https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html) .
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of the place index resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	// - Must be a unique place index resource name.
	// - No spaces allowed. For example, `ExamplePlaceIndex` .
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Specifies the data storage option for requesting Places.
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The optional description for the place index resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
}

