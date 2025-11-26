package previewawscleanroomsmixins


// Provides the input reference configuration for the ID mapping table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingTableInputReferenceConfigProperty := &IdMappingTableInputReferenceConfigProperty{
//   	InputReferenceArn: jsii.String("inputReferenceArn"),
//   	ManageResourcePolicies: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputreferenceconfig.html
//
type CfnIdMappingTablePropsMixin_IdMappingTableInputReferenceConfigProperty struct {
	// The Amazon Resource Name (ARN) of the referenced resource in AWS Entity Resolution .
	//
	// Valid values are ID mapping workflow ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputreferenceconfig.html#cfn-cleanrooms-idmappingtable-idmappingtableinputreferenceconfig-inputreferencearn
	//
	InputReferenceArn *string `field:"optional" json:"inputReferenceArn" yaml:"inputReferenceArn"`
	// When `TRUE` , AWS Clean Rooms manages permissions for the ID mapping table resource.
	//
	// When `FALSE` , the resource owner manages permissions for the ID mapping table resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputreferenceconfig.html#cfn-cleanrooms-idmappingtable-idmappingtableinputreferenceconfig-manageresourcepolicies
	//
	ManageResourcePolicies interface{} `field:"optional" json:"manageResourcePolicies" yaml:"manageResourcePolicies"`
}

