package awsrds


// A version of an engine - for either a cluster, or instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   engineVersion := &engineVersion{
//   	majorVersion: jsii.String("majorVersion"),
//
//   	// the properties below are optional
//   	fullVersion: jsii.String("fullVersion"),
//   }
//
type EngineVersion struct {
	// The major version of the engine, for example, "5.6". Used in specifying the ParameterGroup family and OptionGroup version for this engine.
	MajorVersion *string `field:"required" json:"majorVersion" yaml:"majorVersion"`
	// The full version string of the engine, for example, "5.6.mysql_aurora.1.22.1". It can be undefined, which means RDS should use whatever version it deems appropriate for the given engine type.
	FullVersion *string `field:"optional" json:"fullVersion" yaml:"fullVersion"`
}

