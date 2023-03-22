package awsappflow


// The properties that are applied when using SAPOData as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataSourcePropertiesProperty := &sAPODataSourcePropertiesProperty{
//   	objectPath: jsii.String("objectPath"),
//   }
//
type CfnFlow_SAPODataSourcePropertiesProperty struct {
	// The object path specified in the SAPOData flow source.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
}

