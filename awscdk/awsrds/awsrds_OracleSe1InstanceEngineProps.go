package awsrds


// Properties for Oracle Standard Edition 1 instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleSe1}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var oracleLegacyEngineVersion oracleLegacyEngineVersion
//
//   oracleSe1InstanceEngineProps := &oracleSe1InstanceEngineProps{
//   	version: oracleLegacyEngineVersion,
//   }
//
// Deprecated: instances can no longer be created with this engine. See https://forums.aws.amazon.com/ann.jspa?annID=7341
type OracleSe1InstanceEngineProps struct {
	// The exact version of the engine to use.
	// Deprecated: instances can no longer be created with this engine. See https://forums.aws.amazon.com/ann.jspa?annID=7341
	Version OracleLegacyEngineVersion `field:"required" json:"version" yaml:"version"`
}

