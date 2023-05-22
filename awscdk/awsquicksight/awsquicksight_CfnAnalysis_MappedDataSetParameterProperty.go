package awsquicksight


// A dataset parameter that is mapped to an analysis parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappedDataSetParameterProperty := &MappedDataSetParameterProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	DataSetParameterName: jsii.String("dataSetParameterName"),
//   }
//
type CfnAnalysis_MappedDataSetParameterProperty struct {
	// A unique name that identifies a dataset within the analysis or dashboard.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The name of the dataset parameter.
	DataSetParameterName *string `field:"required" json:"dataSetParameterName" yaml:"dataSetParameterName"`
}

