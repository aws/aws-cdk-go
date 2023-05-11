package awsathena


// The Athena engine version for running queries, or the PySpark engine version for running sessions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   engineVersionProperty := &EngineVersionProperty{
//   	EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   	SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   }
//
type CfnWorkGroup_EngineVersionProperty struct {
	// Read only.
	//
	// The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a `CreateWorkGroup` or `UpdateWorkGroup` operation, the `EffectiveEngineVersion` field is ignored.
	EffectiveEngineVersion *string `field:"optional" json:"effectiveEngineVersion" yaml:"effectiveEngineVersion"`
	// The engine version requested by the user.
	//
	// Possible values are determined by the output of `ListEngineVersions` , including AUTO. The default is AUTO.
	SelectedEngineVersion *string `field:"optional" json:"selectedEngineVersion" yaml:"selectedEngineVersion"`
}

