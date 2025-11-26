package previewawscleanroomsmixins


// Provides the information for the ID namespace association input reference configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idNamespaceAssociationInputReferenceConfigProperty := &IdNamespaceAssociationInputReferenceConfigProperty{
//   	InputReferenceArn: jsii.String("inputReferenceArn"),
//   	ManageResourcePolicies: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig.html
//
type CfnIdNamespaceAssociationPropsMixin_IdNamespaceAssociationInputReferenceConfigProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Entity Resolution resource that is being associated to the collaboration.
	//
	// Valid resource ARNs are from the ID namespaces that you own.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig.html#cfn-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig-inputreferencearn
	//
	InputReferenceArn *string `field:"optional" json:"inputReferenceArn" yaml:"inputReferenceArn"`
	// When `TRUE` , AWS Clean Rooms manages permissions for the ID namespace association resource.
	//
	// When `FALSE` , the resource owner manages permissions for the ID namespace association resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig.html#cfn-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceconfig-manageresourcepolicies
	//
	ManageResourcePolicies interface{} `field:"optional" json:"manageResourcePolicies" yaml:"manageResourcePolicies"`
}

