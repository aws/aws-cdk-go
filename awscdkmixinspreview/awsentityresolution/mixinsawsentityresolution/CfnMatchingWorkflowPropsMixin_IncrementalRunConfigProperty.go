package mixinsawsentityresolution


// Optional.
//
// An object that defines the incremental run type. This object contains only the `incrementalRunType` field, which appears as "Automatic" in the console.
//
// > For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER` , incremental processing is not supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   incrementalRunConfigProperty := &IncrementalRunConfigProperty{
//   	IncrementalRunType: jsii.String("incrementalRunType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-incrementalrunconfig.html
//
type CfnMatchingWorkflowPropsMixin_IncrementalRunConfigProperty struct {
	// The type of incremental run. The only valid value is `IMMEDIATE` . This appears as "Automatic" in the console.
	//
	// > For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER` , incremental processing is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-incrementalrunconfig.html#cfn-entityresolution-matchingworkflow-incrementalrunconfig-incrementalruntype
	//
	IncrementalRunType *string `field:"optional" json:"incrementalRunType" yaml:"incrementalRunType"`
}

