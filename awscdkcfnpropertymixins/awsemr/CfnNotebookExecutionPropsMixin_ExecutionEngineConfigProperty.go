package awsemr


// Specifies the execution engine (cluster) to run the notebook and perform the notebook execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   executionEngineConfigProperty := &ExecutionEngineConfigProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-executionengineconfig.html
//
type CfnNotebookExecutionPropsMixin_ExecutionEngineConfigProperty struct {
	// The unique identifier of the execution engine.
	//
	// For an Amazon EMR cluster, this is the cluster ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-executionengineconfig.html#cfn-emr-notebookexecution-executionengineconfig-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of execution engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-executionengineconfig.html#cfn-emr-notebookexecution-executionengineconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

