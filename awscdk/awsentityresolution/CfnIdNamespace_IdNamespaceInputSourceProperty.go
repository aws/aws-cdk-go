package awsentityresolution


// An object containing `InputSourceARN` and `SchemaName` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idNamespaceInputSourceProperty := &IdNamespaceInputSourceProperty{
//   	InputSourceArn: jsii.String("inputSourceArn"),
//
//   	// the properties below are optional
//   	SchemaName: jsii.String("schemaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html
//
type CfnIdNamespace_IdNamespaceInputSourceProperty struct {
	// An AWS Glue table ARN for the input source table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-inputsourcearn
	//
	InputSourceArn *string `field:"required" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html#cfn-entityresolution-idnamespace-idnamespaceinputsource-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

