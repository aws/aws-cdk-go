package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databricksParametersProperty := &DatabricksParametersProperty{
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   	SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   }
//
type CfnDataSource_DatabricksParametersProperty struct {
	// `CfnDataSource.DatabricksParametersProperty.Host`.
	Host *string `field:"required" json:"host" yaml:"host"`
	// `CfnDataSource.DatabricksParametersProperty.Port`.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// `CfnDataSource.DatabricksParametersProperty.SqlEndpointPath`.
	SqlEndpointPath *string `field:"required" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

