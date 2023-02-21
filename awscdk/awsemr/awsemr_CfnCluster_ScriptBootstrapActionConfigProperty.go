package awsemr


// `ScriptBootstrapActionConfig` is a subproperty of the `BootstrapActionConfig` property type.
//
// `ScriptBootstrapActionConfig` specifies the arguments and location of the bootstrap script for EMR to run on all cluster nodes before it installs open-source big data applications on them.
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
type CfnCluster_ScriptBootstrapActionConfigProperty struct {
	// Location in Amazon S3 of the script to run during a bootstrap action.
	Path *string `field:"required" json:"path" yaml:"path"`
	// A list of command line arguments to pass to the bootstrap action script.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

