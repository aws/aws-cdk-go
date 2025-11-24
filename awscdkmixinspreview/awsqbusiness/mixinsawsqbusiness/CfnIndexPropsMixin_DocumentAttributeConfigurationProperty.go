package mixinsawsqbusiness


// Configuration information for document attributes.
//
// Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
//
// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentAttributeConfigurationProperty := &DocumentAttributeConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Search: jsii.String("search"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-documentattributeconfiguration.html
//
type CfnIndexPropsMixin_DocumentAttributeConfigurationProperty struct {
	// The name of the document attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-documentattributeconfiguration.html#cfn-qbusiness-index-documentattributeconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about whether the document attribute can be used by an end user to search for information on their web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-documentattributeconfiguration.html#cfn-qbusiness-index-documentattributeconfiguration-search
	//
	Search *string `field:"optional" json:"search" yaml:"search"`
	// The type of document attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-documentattributeconfiguration.html#cfn-qbusiness-index-documentattributeconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

