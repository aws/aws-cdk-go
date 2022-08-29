package awsrds


// Properties for Oracle Standard Edition 2 (CDB) instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleSe2Cdb}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var oracleEngineVersion oracleEngineVersion
//
//   oracleSe2CdbInstanceEngineProps := &oracleSe2CdbInstanceEngineProps{
//   	version: oracleEngineVersion,
//   }
//
type OracleSe2CdbInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version OracleEngineVersion `field:"required" json:"version" yaml:"version"`
}

