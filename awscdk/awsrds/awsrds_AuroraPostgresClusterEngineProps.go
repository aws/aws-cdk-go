package awsrds


// Creation properties of the Aurora PostgreSQL database cluster engine.
//
// Used in {@link DatabaseClusterEngine.auroraPostgres}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var auroraPostgresEngineVersion auroraPostgresEngineVersion
//
//   auroraPostgresClusterEngineProps := &auroraPostgresClusterEngineProps{
//   	version: auroraPostgresEngineVersion,
//   }
//
type AuroraPostgresClusterEngineProps struct {
	// The version of the Aurora PostgreSQL cluster engine.
	Version AuroraPostgresEngineVersion `field:"required" json:"version" yaml:"version"`
}

