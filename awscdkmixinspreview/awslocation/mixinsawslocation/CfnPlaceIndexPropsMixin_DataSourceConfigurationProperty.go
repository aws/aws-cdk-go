package mixinsawslocation


// Specifies the data storage option requesting Places.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	IntendedUse: jsii.String("intendedUse"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-placeindex-datasourceconfiguration.html
//
type CfnPlaceIndexPropsMixin_DataSourceConfigurationProperty struct {
	// Specifies how the results of an operation will be stored by the caller.
	//
	// Valid values include:
	//
	// - `SingleUse` specifies that the results won't be stored.
	// - `Storage` specifies that the result can be cached or stored in a database.
	//
	// Default value: `SingleUse`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-placeindex-datasourceconfiguration.html#cfn-location-placeindex-datasourceconfiguration-intendeduse
	//
	IntendedUse *string `field:"optional" json:"intendedUse" yaml:"intendedUse"`
}

