package previewawsdataexchangeevents


// Type definition for DataSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSet := &DataSet{
//   	AssetType: []*string{
//   		jsii.String("assetType"),
//   	},
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type DeprecationPlannedForDataSet_DataSet struct {
	// AssetType property.
	//
	// Specify an array of string values to match this event if the actual value of AssetType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AssetType *[]*string `field:"optional" json:"assetType" yaml:"assetType"`
	// Id property.
	//
	// Specify an array of string values to match this event if the actual value of Id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// Name property.
	//
	// Specify an array of string values to match this event if the actual value of Name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

