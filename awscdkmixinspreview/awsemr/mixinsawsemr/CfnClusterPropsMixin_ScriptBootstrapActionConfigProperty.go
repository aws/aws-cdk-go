package mixinsawsemr


// `ScriptBootstrapActionConfig` is a subproperty of the `BootstrapActionConfig` property type.
//
// `ScriptBootstrapActionConfig` specifies the arguments and location of the bootstrap script for EMR to run on all cluster nodes before it installs open-source big data applications on them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scriptBootstrapActionConfigProperty := &ScriptBootstrapActionConfigProperty{
//   	Args: []*string{
//   		jsii.String("args"),
//   	},
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scriptbootstrapactionconfig.html
//
type CfnClusterPropsMixin_ScriptBootstrapActionConfigProperty struct {
	// A list of command line arguments to pass to the bootstrap action script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scriptbootstrapactionconfig.html#cfn-emr-cluster-scriptbootstrapactionconfig-args
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Location in Amazon S3 of the script to run during a bootstrap action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-scriptbootstrapactionconfig.html#cfn-emr-cluster-scriptbootstrapactionconfig-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

