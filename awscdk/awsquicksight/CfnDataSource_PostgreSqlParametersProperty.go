package awsquicksight


// The parameters for PostgreSQL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postgreSqlParametersProperty := &PostgreSqlParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-postgresqlparameters.html
//
type CfnDataSource_PostgreSqlParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-postgresqlparameters.html#cfn-quicksight-datasource-postgresqlparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-postgresqlparameters.html#cfn-quicksight-datasource-postgresqlparameters-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-postgresqlparameters.html#cfn-quicksight-datasource-postgresqlparameters-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

