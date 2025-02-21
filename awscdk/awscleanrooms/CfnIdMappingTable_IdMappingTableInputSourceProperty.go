package awscleanrooms


// The input source of the ID mapping table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingTableInputSourceProperty := &IdMappingTableInputSourceProperty{
//   	IdNamespaceAssociationId: jsii.String("idNamespaceAssociationId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputsource.html
//
type CfnIdMappingTable_IdMappingTableInputSourceProperty struct {
	// The unique identifier of the ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputsource.html#cfn-cleanrooms-idmappingtable-idmappingtableinputsource-idnamespaceassociationid
	//
	IdNamespaceAssociationId *string `field:"required" json:"idNamespaceAssociationId" yaml:"idNamespaceAssociationId"`
	// The type of the input source of the ID mapping table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputsource.html#cfn-cleanrooms-idmappingtable-idmappingtableinputsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

