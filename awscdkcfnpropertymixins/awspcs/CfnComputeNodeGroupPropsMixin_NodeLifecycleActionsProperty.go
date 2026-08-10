package awspcs


// Custom scripts that run at defined points in a compute node's lifecycle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nodeLifecycleActionsProperty := &NodeLifecycleActionsProperty{
//   	ScriptCachingPolicy: jsii.String("scriptCachingPolicy"),
//   	Stages: &NodeLifecycleStagesProperty{
//   		NodeBootstrapped: []interface{}{
//   			&NodeLifecycleScriptProperty{
//   				Arguments: []*string{
//   					jsii.String("arguments"),
//   				},
//   				ExecutionPolicy: jsii.String("executionPolicy"),
//   				Name: jsii.String("name"),
//   				OnError: jsii.String("onError"),
//   				ScriptSource: &ScriptSourceProperty{
//   					Checksum: jsii.String("checksum"),
//   					S3VersionId: jsii.String("s3VersionId"),
//   					ScriptLocation: jsii.String("scriptLocation"),
//   				},
//   			},
//   		},
//   		NodeReady: []interface{}{
//   			&NodeLifecycleScriptProperty{
//   				Arguments: []*string{
//   					jsii.String("arguments"),
//   				},
//   				ExecutionPolicy: jsii.String("executionPolicy"),
//   				Name: jsii.String("name"),
//   				OnError: jsii.String("onError"),
//   				ScriptSource: &ScriptSourceProperty{
//   					Checksum: jsii.String("checksum"),
//   					S3VersionId: jsii.String("s3VersionId"),
//   					ScriptLocation: jsii.String("scriptLocation"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html
//
type CfnComputeNodeGroupPropsMixin_NodeLifecycleActionsProperty struct {
	// Controls whether lifecycle scripts are downloaded once at first boot (CACHE_ONCE) or re-downloaded on every reboot (REFRESH_ON_REBOOT).
	//
	// Defaults to CACHE_ONCE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html#cfn-pcs-computenodegroup-nodelifecycleactions-scriptcachingpolicy
	//
	ScriptCachingPolicy *string `field:"optional" json:"scriptCachingPolicy" yaml:"scriptCachingPolicy"`
	// The ordered scripts to run at each compute node lifecycle stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html#cfn-pcs-computenodegroup-nodelifecycleactions-stages
	//
	Stages interface{} `field:"optional" json:"stages" yaml:"stages"`
}

