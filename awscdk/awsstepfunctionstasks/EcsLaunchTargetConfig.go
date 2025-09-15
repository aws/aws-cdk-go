package awsstepfunctionstasks


// Configuration options for the ECS launch type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   ecsLaunchTargetConfig := &EcsLaunchTargetConfig{
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   }
//
type EcsLaunchTargetConfig struct {
	// Additional parameters to pass to the base task.
	// Default: - No additional parameters passed.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

