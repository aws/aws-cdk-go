package previewawscleanroomsmixins


// The configuration settings for the ID mapping table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingConfigProperty := &IdMappingConfigProperty{
//   	AllowUseAsDimensionColumn: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idmappingconfig.html
//
type CfnIdNamespaceAssociationPropsMixin_IdMappingConfigProperty struct {
	// An indicator as to whether you can use your column as a dimension column in the ID mapping table ( `TRUE` ) or not ( `FALSE` ).
	//
	// Default is `FALSE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-idnamespaceassociation-idmappingconfig.html#cfn-cleanrooms-idnamespaceassociation-idmappingconfig-allowuseasdimensioncolumn
	//
	AllowUseAsDimensionColumn interface{} `field:"optional" json:"allowUseAsDimensionColumn" yaml:"allowUseAsDimensionColumn"`
}

