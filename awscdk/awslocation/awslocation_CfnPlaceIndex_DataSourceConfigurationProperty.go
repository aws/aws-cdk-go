package awslocation


// Specifies the data storage option for requesting Places.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &dataSourceConfigurationProperty{
//   	intendedUse: jsii.String("intendedUse"),
//   }
//
type CfnPlaceIndex_DataSourceConfigurationProperty struct {
	// Specifies how the results of an operation will be stored by the caller.
	//
	// Valid values include:
	//
	// - `SingleUse` specifies that the results won't be stored.
	// - `Storage` specifies that the result can be cached or stored in a database.
	//
	// > Place index resources using HERE as a data provider can't be configured to store results for locations in Japan when choosing `Storage` for the `IntendedUse` parameter.
	//
	// Default value: `SingleUse`.
	IntendedUse *string `field:"optional" json:"intendedUse" yaml:"intendedUse"`
}

