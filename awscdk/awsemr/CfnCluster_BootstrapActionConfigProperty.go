package awsemr


// `BootstrapActionConfig` is a property of `AWS::EMR::Cluster` that can be used to run bootstrap actions on EMR clusters.
//
// You can use a bootstrap action to install software and configure EC2 instances for all cluster nodes before EMR installs and configures open-source big data applications on cluster instances. For more information, see [Create Bootstrap Actions to Install Additional Software](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-plan-bootstrap.html) in the *Amazon EMR Management Guide* .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html
//
type CfnCluster_BootstrapActionConfigProperty struct {
	// The name of the bootstrap action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The script run by the bootstrap action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction
	//
	ScriptBootstrapAction interface{} `field:"required" json:"scriptBootstrapAction" yaml:"scriptBootstrapAction"`
}

