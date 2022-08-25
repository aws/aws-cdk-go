package awsstepfunctionstasks


// Configuration of the script to run during a bootstrap action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptBootstrapActionConfigProperty := &scriptBootstrapActionConfigProperty{
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScriptBootstrapActionConfig.html
//
type EmrCreateCluster_ScriptBootstrapActionConfigProperty struct {
	// Location of the script to run during a bootstrap action.
	//
	// Can be either a location in Amazon S3 or on a local file system.
	Path *string `field:"required" json:"path" yaml:"path"`
	// A list of command line arguments to pass to the bootstrap action script.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

