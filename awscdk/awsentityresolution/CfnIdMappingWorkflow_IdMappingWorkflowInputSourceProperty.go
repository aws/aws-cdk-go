package awsentityresolution


// An object containing `InputSourceARN` and `SchemaName` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingWorkflowInputSourceProperty := &IdMappingWorkflowInputSourceProperty{
//   	InputSourceArn: jsii.String("inputSourceArn"),
//   	SchemaArn: jsii.String("schemaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html
//
type CfnIdMappingWorkflow_IdMappingWorkflowInputSourceProperty struct {
	// An AWS Glue table ARN for the input source table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowinputsource-inputsourcearn
	//
	InputSourceArn *string `field:"required" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The ARN (Amazon Resource Name) that AWS Entity Resolution generated for the `SchemaMapping` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowinputsource-schemaarn
	//
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
}

