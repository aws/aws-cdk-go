package awslocation


// Specifies the data storage option requesting Places.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	IntendedUse: jsii.String("intendedUse"),
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
	// Default value: `SingleUse`.
	IntendedUse *string `field:"optional" json:"intendedUse" yaml:"intendedUse"`
}

