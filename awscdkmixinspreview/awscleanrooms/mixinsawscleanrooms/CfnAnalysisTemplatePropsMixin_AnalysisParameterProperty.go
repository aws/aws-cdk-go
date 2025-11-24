package mixinsawscleanrooms


// Optional.
//
// The member who can query can provide this placeholder for a literal data value in an analysis template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisParameterProperty := &AnalysisParameterProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisparameter.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisParameterProperty struct {
	// Optional.
	//
	// The default value that is applied in the analysis template. The member who can query can override this value in the query editor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisparameter.html#cfn-cleanrooms-analysistemplate-analysisparameter-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The name of the parameter.
	//
	// The name must use only alphanumeric, underscore (_), or hyphen (-) characters but cannot start or end with a hyphen.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisparameter.html#cfn-cleanrooms-analysistemplate-analysisparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisparameter.html#cfn-cleanrooms-analysistemplate-analysisparameter-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

