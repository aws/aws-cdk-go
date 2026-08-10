package awspcs


// A single lifecycle script with its source, arguments, and error behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nodeLifecycleScriptProperty := &NodeLifecycleScriptProperty{
//   	Arguments: []*string{
//   		jsii.String("arguments"),
//   	},
//   	ExecutionPolicy: jsii.String("executionPolicy"),
//   	Name: jsii.String("name"),
//   	OnError: jsii.String("onError"),
//   	ScriptSource: &ScriptSourceProperty{
//   		Checksum: jsii.String("checksum"),
//   		S3VersionId: jsii.String("s3VersionId"),
//   		ScriptLocation: jsii.String("scriptLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html
//
type CfnComputeNodeGroupPropsMixin_NodeLifecycleScriptProperty struct {
	// An ordered list of arguments passed to the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html#cfn-pcs-computenodegroup-nodelifecyclescript-arguments
	//
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Whether the script runs only on the node's first boot (FIRST_BOOT_ONLY) or on every boot including reboots (EVERY_BOOT).
	//
	// Defaults to FIRST_BOOT_ONLY.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html#cfn-pcs-computenodegroup-nodelifecyclescript-executionpolicy
	//
	ExecutionPolicy *string `field:"optional" json:"executionPolicy" yaml:"executionPolicy"`
	// A human-readable name that identifies the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html#cfn-pcs-computenodegroup-nodelifecyclescript-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The behavior when the script exits with an error.
	//
	// Defaults to TERMINATE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html#cfn-pcs-computenodegroup-nodelifecyclescript-onerror
	//
	OnError *string `field:"optional" json:"onError" yaml:"onError"`
	// The external location of a lifecycle script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclescript.html#cfn-pcs-computenodegroup-nodelifecyclescript-scriptsource
	//
	ScriptSource interface{} `field:"optional" json:"scriptSource" yaml:"scriptSource"`
}

