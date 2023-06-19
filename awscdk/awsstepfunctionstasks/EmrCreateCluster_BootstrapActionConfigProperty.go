package awsstepfunctionstasks


// Configuration of a bootstrap action.
//
// See the RunJobFlow API for complete documentation on input parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstrapActionConfigProperty := &BootstrapActionConfigProperty{
//   	Name: jsii.String("name"),
//   	ScriptBootstrapAction: &ScriptBootstrapActionConfigProperty{
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		Args: []*string{
//   			jsii.String("args"),
//   		},
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_BootstrapActionConfig.html
//
// Experimental.
type EmrCreateCluster_BootstrapActionConfigProperty struct {
	// The name of the bootstrap action.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The script run by the bootstrap action.
	// Experimental.
	ScriptBootstrapAction *EmrCreateCluster_ScriptBootstrapActionConfigProperty `field:"required" json:"scriptBootstrapAction" yaml:"scriptBootstrapAction"`
}

