package previewawscleanroomsmixins


// The input reference properties for the ID mapping table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingTableInputReferencePropertiesProperty := &IdMappingTableInputReferencePropertiesProperty{
//   	IdMappingTableInputSource: []interface{}{
//   		&IdMappingTableInputSourceProperty{
//   			IdNamespaceAssociationId: jsii.String("idNamespaceAssociationId"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputreferenceproperties.html
//
type CfnIdMappingTablePropsMixin_IdMappingTableInputReferencePropertiesProperty struct {
	// The input source of the ID mapping table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idmappingtable-idmappingtableinputreferenceproperties.html#cfn-cleanrooms-idmappingtable-idmappingtableinputreferenceproperties-idmappingtableinputsource
	//
	IdMappingTableInputSource interface{} `field:"optional" json:"idMappingTableInputSource" yaml:"idMappingTableInputSource"`
}

