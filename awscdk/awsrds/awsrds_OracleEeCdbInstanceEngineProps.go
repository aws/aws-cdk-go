package awsrds


// Properties for Oracle Enterprise Edition (CDB) instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleEeCdb}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var oracleEngineVersion oracleEngineVersion
//
//   oracleEeCdbInstanceEngineProps := &oracleEeCdbInstanceEngineProps{
//   	version: oracleEngineVersion,
//   }
//
type OracleEeCdbInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `field:"required" json:"version" yaml:"version"`
}

