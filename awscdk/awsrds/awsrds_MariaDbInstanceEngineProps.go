package awsrds


// Properties for MariaDB instance engines.
//
// Used in {@link DatabaseInstanceEngine.mariaDb}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mariaDbEngineVersion mariaDbEngineVersion
//
//   mariaDbInstanceEngineProps := &mariaDbInstanceEngineProps{
//   	version: mariaDbEngineVersion,
//   }
//
type MariaDbInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version MariaDbEngineVersion `field:"required" json:"version" yaml:"version"`
}

