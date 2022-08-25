package awsrds


// Properties for SQL Server Standard Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerSe}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sqlServerEngineVersion sqlServerEngineVersion
//
//   sqlServerSeInstanceEngineProps := &sqlServerSeInstanceEngineProps{
//   	version: sqlServerEngineVersion,
//   }
//
type SqlServerSeInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version SqlServerEngineVersion `field:"required" json:"version" yaml:"version"`
}

