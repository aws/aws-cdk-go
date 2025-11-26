package previewawsquicksightmixins


// Parameters for Amazon Aurora PostgreSQL-Compatible Edition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   auroraPostgreSqlParametersProperty := &AuroraPostgreSqlParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-aurorapostgresqlparameters.html
//
type CfnDataSourcePropsMixin_AuroraPostgreSqlParametersProperty struct {
	// The Amazon Aurora PostgreSQL database to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-aurorapostgresqlparameters.html#cfn-quicksight-datasource-aurorapostgresqlparameters-database
	//
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The Amazon Aurora PostgreSQL-Compatible host to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-aurorapostgresqlparameters.html#cfn-quicksight-datasource-aurorapostgresqlparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The port that Amazon Aurora PostgreSQL is listening on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-aurorapostgresqlparameters.html#cfn-quicksight-datasource-aurorapostgresqlparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

