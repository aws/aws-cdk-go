package awsentityresolution


// An object containing `inputSourceARN` , `schemaName` , and `applyNormalization` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSourceProperty := &InputSourceProperty{
//   	InputSourceArn: jsii.String("inputSourceArn"),
//   	SchemaArn: jsii.String("schemaArn"),
//
//   	// the properties below are optional
//   	ApplyNormalization: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html
//
type CfnMatchingWorkflow_InputSourceProperty struct {
	// An object containing `inputSourceARN` , `schemaName` , and `applyNormalization` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-inputsourcearn
	//
	InputSourceArn *string `field:"required" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-schemaarn
	//
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
	// Normalizes the attributes defined in the schema in the input data.
	//
	// For example, if an attribute has an `AttributeType` of `PHONE_NUMBER` , and the data in the input table is in a format of 1234567890, AWS Entity Resolution will normalize this field in the output to (123)-456-7890.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-applynormalization
	//
	ApplyNormalization interface{} `field:"optional" json:"applyNormalization" yaml:"applyNormalization"`
}

