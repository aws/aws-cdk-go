package mixinsawsentityresolution


// An object containing `inputSourceARN` , `schemaName` , and `applyNormalization` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputSourceProperty := &InputSourceProperty{
//   	ApplyNormalization: jsii.Boolean(false),
//   	InputSourceArn: jsii.String("inputSourceArn"),
//   	SchemaArn: jsii.String("schemaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html
//
type CfnMatchingWorkflowPropsMixin_InputSourceProperty struct {
	// Normalizes the attributes defined in the schema in the input data.
	//
	// For example, if an attribute has an `AttributeType` of `PHONE_NUMBER` , and the data in the input table is in a format of 1234567890, AWS Entity Resolution will normalize this field in the output to (123)-456-7890.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-applynormalization
	//
	ApplyNormalization interface{} `field:"optional" json:"applyNormalization" yaml:"applyNormalization"`
	// An object containing `inputSourceARN` , `schemaName` , and `applyNormalization` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-inputsourcearn
	//
	InputSourceArn *string `field:"optional" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-schemaarn
	//
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
}

