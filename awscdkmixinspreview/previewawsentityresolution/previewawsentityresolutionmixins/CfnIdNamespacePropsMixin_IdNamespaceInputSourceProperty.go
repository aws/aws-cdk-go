package previewawsentityresolutionmixins


// An object containing `inputSourceARN` and `schemaName` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idNamespaceInputSourceProperty := &IdNamespaceInputSourceProperty{
//   	InputSourceArn: jsii.String("inputSourceArn"),
//   	SchemaName: jsii.String("schemaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html
//
type CfnIdNamespacePropsMixin_IdNamespaceInputSourceProperty struct {
	// An AWS Glue table Amazon Resource Name (ARN) or a matching workflow ARN for the input source table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-inputsourcearn
	//
	InputSourceArn *string `field:"optional" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

