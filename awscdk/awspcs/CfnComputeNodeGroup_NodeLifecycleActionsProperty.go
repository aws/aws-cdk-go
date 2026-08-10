package awspcs


// Custom scripts that run at defined points in a compute node's lifecycle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeLifecycleActionsProperty := &NodeLifecycleActionsProperty{
//   	Stages: &NodeLifecycleStagesProperty{
//   		NodeBootstrapped: []interface{}{
//   			&NodeLifecycleScriptProperty{
//   				Name: jsii.String("name"),
//   				ScriptSource: &ScriptSourceProperty{
//   					ScriptLocation: jsii.String("scriptLocation"),
//
//   					// the properties below are optional
//   					Checksum: jsii.String("checksum"),
//   					S3VersionId: jsii.String("s3VersionId"),
//   				},
//
//   				// the properties below are optional
//   				Arguments: []*string{
//   					jsii.String("arguments"),
//   				},
//   				ExecutionPolicy: jsii.String("executionPolicy"),
//   				OnError: jsii.String("onError"),
//   			},
//   		},
//   		NodeReady: []interface{}{
//   			&NodeLifecycleScriptProperty{
//   				Name: jsii.String("name"),
//   				ScriptSource: &ScriptSourceProperty{
//   					ScriptLocation: jsii.String("scriptLocation"),
//
//   					// the properties below are optional
//   					Checksum: jsii.String("checksum"),
//   					S3VersionId: jsii.String("s3VersionId"),
//   				},
//
//   				// the properties below are optional
//   				Arguments: []*string{
//   					jsii.String("arguments"),
//   				},
//   				ExecutionPolicy: jsii.String("executionPolicy"),
//   				OnError: jsii.String("onError"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ScriptCachingPolicy: jsii.String("scriptCachingPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html
//
type CfnComputeNodeGroup_NodeLifecycleActionsProperty struct {
	// The ordered scripts to run at each compute node lifecycle stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html#cfn-pcs-computenodegroup-nodelifecycleactions-stages
	//
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// Controls whether lifecycle scripts are downloaded once at first boot (CACHE_ONCE) or re-downloaded on every reboot (REFRESH_ON_REBOOT).
	//
	// Defaults to CACHE_ONCE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-nodelifecycleactions.html#cfn-pcs-computenodegroup-nodelifecycleactions-scriptcachingpolicy
	//
	ScriptCachingPolicy *string `field:"optional" json:"scriptCachingPolicy" yaml:"scriptCachingPolicy"`
}

