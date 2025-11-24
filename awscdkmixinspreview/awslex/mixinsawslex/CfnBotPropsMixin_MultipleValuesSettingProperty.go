package mixinsawslex


// Indicates whether a slot can return multiple values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multipleValuesSettingProperty := &MultipleValuesSettingProperty{
//   	AllowMultipleValues: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-multiplevaluessetting.html
//
type CfnBotPropsMixin_MultipleValuesSettingProperty struct {
	// Indicates whether a slot can return multiple values.
	//
	// When `true` , the slot may return more than one value in a response. When `false` , the slot returns only a single value.
	//
	// Multi-value slots are only available in the en-US locale. If you set this value to `true` in any other locale, Amazon Lex throws a `ValidationException` .
	//
	// If the `allowMutlipleValues` is not set, the default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-multiplevaluessetting.html#cfn-lex-bot-multiplevaluessetting-allowmultiplevalues
	//
	AllowMultipleValues interface{} `field:"optional" json:"allowMultipleValues" yaml:"allowMultipleValues"`
}

