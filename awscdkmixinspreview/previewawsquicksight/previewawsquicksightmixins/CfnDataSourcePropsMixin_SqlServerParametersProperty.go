package previewawsquicksightmixins


// The parameters for SQL Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sqlServerParametersProperty := &SqlServerParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-sqlserverparameters.html
//
type CfnDataSourcePropsMixin_SqlServerParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-sqlserverparameters.html#cfn-quicksight-datasource-sqlserverparameters-database
	//
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-sqlserverparameters.html#cfn-quicksight-datasource-sqlserverparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-sqlserverparameters.html#cfn-quicksight-datasource-sqlserverparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

