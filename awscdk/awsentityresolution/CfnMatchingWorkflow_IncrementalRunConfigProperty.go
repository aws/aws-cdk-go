package awsentityresolution


// An object which defines an incremental run type and has only `incrementalRunType` as a field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incrementalRunConfigProperty := &IncrementalRunConfigProperty{
//   	IncrementalRunType: jsii.String("incrementalRunType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-incrementalrunconfig.html
//
type CfnMatchingWorkflow_IncrementalRunConfigProperty struct {
	// The type of incremental run.
	//
	// It takes only one value: `IMMEDIATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-incrementalrunconfig.html#cfn-entityresolution-matchingworkflow-incrementalrunconfig-incrementalruntype
	//
	IncrementalRunType *string `field:"required" json:"incrementalRunType" yaml:"incrementalRunType"`
}

