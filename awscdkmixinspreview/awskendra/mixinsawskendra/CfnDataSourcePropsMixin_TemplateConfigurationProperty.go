package mixinsawskendra


// Provides a template for the configuration information to connect to your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateConfigurationProperty := &TemplateConfigurationProperty{
//   	Template: jsii.String("template"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-templateconfiguration.html
//
type CfnDataSourcePropsMixin_TemplateConfigurationProperty struct {
	// The template schema used for the data source, where templates schemas are supported.
	//
	// See [Data source template schemas](https://docs.aws.amazon.com/kendra/latest/dg/ds-schemas.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-templateconfiguration.html#cfn-kendra-datasource-templateconfiguration-template
	//
	Template *string `field:"optional" json:"template" yaml:"template"`
}

