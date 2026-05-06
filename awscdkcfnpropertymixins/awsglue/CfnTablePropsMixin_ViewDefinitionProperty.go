package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   viewDefinitionProperty := &ViewDefinitionProperty{
//   	Definer: jsii.String("definer"),
//   	IsProtected: jsii.Boolean(false),
//   	Representations: []interface{}{
//   		&ViewRepresentationProperty{
//   			Dialect: jsii.String("dialect"),
//   			DialectVersion: jsii.String("dialectVersion"),
//   			ValidationConnection: jsii.String("validationConnection"),
//   			ViewExpandedText: jsii.String("viewExpandedText"),
//   			ViewOriginalText: jsii.String("viewOriginalText"),
//   		},
//   	},
//   	SubObjects: []*string{
//   		jsii.String("subObjects"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewdefinition.html
//
type CfnTablePropsMixin_ViewDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewdefinition.html#cfn-glue-table-viewdefinition-definer
	//
	Definer *string `field:"optional" json:"definer" yaml:"definer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewdefinition.html#cfn-glue-table-viewdefinition-isprotected
	//
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewdefinition.html#cfn-glue-table-viewdefinition-representations
	//
	Representations interface{} `field:"optional" json:"representations" yaml:"representations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewdefinition.html#cfn-glue-table-viewdefinition-subobjects
	//
	SubObjects *[]*string `field:"optional" json:"subObjects" yaml:"subObjects"`
}

