package awscleanrooms


// Provides the information for the ID namespace association input reference properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var idMappingWorkflowsSupported interface{}
//
//   idNamespaceAssociationInputReferencePropertiesProperty := &IdNamespaceAssociationInputReferencePropertiesProperty{
//   	IdMappingWorkflowsSupported: []interface{}{
//   		idMappingWorkflowsSupported,
//   	},
//   	IdNamespaceType: jsii.String("idNamespaceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceproperties.html
//
type CfnIdNamespaceAssociation_IdNamespaceAssociationInputReferencePropertiesProperty struct {
	// Defines how ID mapping workflows are supported for this ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceproperties.html#cfn-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceproperties-idmappingworkflowssupported
	//
	IdMappingWorkflowsSupported interface{} `field:"optional" json:"idMappingWorkflowsSupported" yaml:"idMappingWorkflowsSupported"`
	// The ID namespace type for this ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceproperties.html#cfn-cleanrooms-idnamespaceassociation-idnamespaceassociationinputreferenceproperties-idnamespacetype
	//
	IdNamespaceType *string `field:"optional" json:"idNamespaceType" yaml:"idNamespaceType"`
}

