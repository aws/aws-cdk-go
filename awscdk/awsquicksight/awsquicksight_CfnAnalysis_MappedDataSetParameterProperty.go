package awsquicksight


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
	// `CfnAnalysis.MappedDataSetParameterProperty.DataSetIdentifier`.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// `CfnAnalysis.MappedDataSetParameterProperty.DataSetParameterName`.
	DataSetParameterName *string `field:"required" json:"dataSetParameterName" yaml:"dataSetParameterName"`
}

