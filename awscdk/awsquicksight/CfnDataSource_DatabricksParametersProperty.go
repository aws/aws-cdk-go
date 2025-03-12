package awsquicksight


// The required parameters that are needed to connect to a Databricks data source.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-databricksparameters.html
//
type CfnDataSource_DatabricksParametersProperty struct {
	// The host name of the Databricks data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-databricksparameters.html#cfn-quicksight-datasource-databricksparameters-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port for the Databricks data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-databricksparameters.html#cfn-quicksight-datasource-databricksparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The HTTP path of the Databricks data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-databricksparameters.html#cfn-quicksight-datasource-databricksparameters-sqlendpointpath
	//
	SqlEndpointPath *string `field:"required" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

