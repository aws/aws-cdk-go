package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingIncrementalRunConfigProperty := &IdMappingIncrementalRunConfigProperty{
//   	IncrementalRunType: jsii.String("incrementalRunType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingincrementalrunconfig.html
//
type CfnIdMappingWorkflowPropsMixin_IdMappingIncrementalRunConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingincrementalrunconfig.html#cfn-entityresolution-idmappingworkflow-idmappingincrementalrunconfig-incrementalruntype
	//
	IncrementalRunType *string `field:"optional" json:"incrementalRunType" yaml:"incrementalRunType"`
}

