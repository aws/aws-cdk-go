package awspcs


// The ordered scripts to run at each compute node lifecycle stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeLifecycleStagesProperty := &NodeLifecycleStagesProperty{
//   	NodeBootstrapped: []interface{}{
//   		&NodeLifecycleScriptProperty{
//   			Name: jsii.String("name"),
//   			ScriptSource: &ScriptSourceProperty{
//   				ScriptLocation: jsii.String("scriptLocation"),
//
//   				// the properties below are optional
//   				Checksum: jsii.String("checksum"),
//   				S3VersionId: jsii.String("s3VersionId"),
//   			},
//
//   			// the properties below are optional
//   			Arguments: []*string{
//   				jsii.String("arguments"),
//   			},
//   			ExecutionPolicy: jsii.String("executionPolicy"),
//   			OnError: jsii.String("onError"),
//   		},
//   	},
//   	NodeReady: []interface{}{
//   		&NodeLifecycleScriptProperty{
//   			Name: jsii.String("name"),
//   			ScriptSource: &ScriptSourceProperty{
//   				ScriptLocation: jsii.String("scriptLocation"),
//
//   				// the properties below are optional
//   				Checksum: jsii.String("checksum"),
//   				S3VersionId: jsii.String("s3VersionId"),
//   			},
//
//   			// the properties below are optional
//   			Arguments: []*string{
//   				jsii.String("arguments"),
//   			},
//   			ExecutionPolicy: jsii.String("executionPolicy"),
//   			OnError: jsii.String("onError"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclestages.html
//
type CfnComputeNodeGroup_NodeLifecycleStagesProperty struct {
	// Scripts to run after the node is bootstrapped, once the PCS configuration phase completes and before slurmd starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclestages.html#cfn-pcs-computenodegroup-nodelifecyclestages-nodebootstrapped
	//
	NodeBootstrapped interface{} `field:"optional" json:"nodeBootstrapped" yaml:"nodeBootstrapped"`
	// Scripts to execute when the node becomes ready (every boot).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecyclestages.html#cfn-pcs-computenodegroup-nodelifecyclestages-nodeready
	//
	NodeReady interface{} `field:"optional" json:"nodeReady" yaml:"nodeReady"`
}

