package awslakeformation


// A reference to a DataLakeSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakeSettingsReference := &DataLakeSettingsReference{
//   	DataLakeSettingsId: jsii.String("dataLakeSettingsId"),
//   }
//
type DataLakeSettingsReference struct {
	// The Id of the DataLakeSettings resource.
	DataLakeSettingsId *string `field:"required" json:"dataLakeSettingsId" yaml:"dataLakeSettingsId"`
}

