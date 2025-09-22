package awsentityresolution


// An object containing `inputSourceARN` , `schemaName` , and `type` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingWorkflowInputSourceProperty := &IdMappingWorkflowInputSourceProperty{
//   	InputSourceArn: jsii.String("inputSourceArn"),
//
//   	// the properties below are optional
//   	SchemaArn: jsii.String("schemaArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html
//
type CfnIdMappingWorkflow_IdMappingWorkflowInputSourceProperty struct {
	// An AWS Glue table Amazon Resource Name (ARN) or a matching workflow ARN for the input source table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowinputsource-inputsourcearn
	//
	InputSourceArn *string `field:"required" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The ARN (Amazon Resource Name) that AWS Entity Resolution generated for the `SchemaMapping` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowinputsource-schemaarn
	//
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
	//
	// The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
	//
	// The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowinputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowinputsource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

