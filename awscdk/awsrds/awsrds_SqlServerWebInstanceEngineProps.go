package awsrds


// Properties for SQL Server Web Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerWeb}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sqlServerEngineVersion sqlServerEngineVersion
//
//   sqlServerWebInstanceEngineProps := &sqlServerWebInstanceEngineProps{
//   	version: sqlServerEngineVersion,
//   }
//
type SqlServerWebInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `field:"required" json:"version" yaml:"version"`
}

