package awsquicksight


// Parameters for Amazon Aurora PostgreSQL-Compatible Edition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraPostgreSqlParametersProperty := &auroraPostgreSqlParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_AuroraPostgreSqlParametersProperty struct {
	// The Amazon Aurora PostgreSQL database to connect to.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The Amazon Aurora PostgreSQL-Compatible host to connect to.
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port that Amazon Aurora PostgreSQL is listening on.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

