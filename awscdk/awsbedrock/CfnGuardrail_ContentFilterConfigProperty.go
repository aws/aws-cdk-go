package awsbedrock


// Contains filter strengths for harmful content.
//
// Guardrails support the following content filters to detect and filter harmful user inputs and FM-generated outputs.
//
// - *Hate* – Describes language or a statement that discriminates, criticizes, insults, denounces, or dehumanizes a person or group on the basis of an identity (such as race, ethnicity, gender, religion, sexual orientation, ability, and national origin).
// - *Insults* – Describes language or a statement that includes demeaning, humiliating, mocking, insulting, or belittling language. This type of language is also labeled as bullying.
// - *Sexual* – Describes language or a statement that indicates sexual interest, activity, or arousal using direct or indirect references to body parts, physical traits, or sex.
// - *Violence* – Describes language or a statement that includes glorification of or threats to inflict physical pain, hurt, or injury toward a person, group or thing.
//
// Content filtering depends on the confidence classification of user inputs and FM responses across each of the four harmful categories. All input and output statements are classified into one of four confidence levels (NONE, LOW, MEDIUM, HIGH) for each harmful category. For example, if a statement is classified as *Hate* with HIGH confidence, the likelihood of the statement representing hateful content is high. A single statement can be classified across multiple categories with varying confidence levels. For example, a single statement can be classified as *Hate* with HIGH confidence, *Insults* with LOW confidence, *Sexual* with NONE confidence, and *Violence* with MEDIUM confidence.
//
// For more information, see [Guardrails content filters](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-filters.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentFilterConfigProperty := &ContentFilterConfigProperty{
//   	InputStrength: jsii.String("inputStrength"),
//   	OutputStrength: jsii.String("outputStrength"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	InputAction: jsii.String("inputAction"),
//   	InputEnabled: jsii.Boolean(false),
//   	InputModalities: []*string{
//   		jsii.String("inputModalities"),
//   	},
//   	OutputAction: jsii.String("outputAction"),
//   	OutputEnabled: jsii.Boolean(false),
//   	OutputModalities: []*string{
//   		jsii.String("outputModalities"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html
//
type CfnGuardrail_ContentFilterConfigProperty struct {
	// The strength of the content filter to apply to prompts.
	//
	// As you increase the filter strength, the likelihood of filtering harmful content increases and the probability of seeing harmful content in your application reduces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-inputstrength
	//
	InputStrength *string `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// The strength of the content filter to apply to model responses.
	//
	// As you increase the filter strength, the likelihood of filtering harmful content increases and the probability of seeing harmful content in your application reduces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-outputstrength
	//
	OutputStrength *string `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// The harmful category that the content filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the action to take when harmful content is detected. Supported values include:.
	//
	// - `BLOCK` – Block the content and replace it with blocked messaging.
	// - `NONE` – Take no action but return detection information in the trace response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-inputaction
	//
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Specifies whether to enable guardrail evaluation on the input.
	//
	// When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-inputenabled
	//
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The input modalities selected for the guardrail content filter configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-inputmodalities
	//
	InputModalities *[]*string `field:"optional" json:"inputModalities" yaml:"inputModalities"`
	// Specifies the action to take when harmful content is detected in the output. Supported values include:.
	//
	// - `BLOCK` – Block the content and replace it with blocked messaging.
	// - `NONE` – Take no action but return detection information in the trace response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-outputaction
	//
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Specifies whether to enable guardrail evaluation on the output.
	//
	// When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-outputenabled
	//
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// The output modalities selected for the guardrail content filter configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-contentfilterconfig.html#cfn-bedrock-guardrail-contentfilterconfig-outputmodalities
	//
	OutputModalities *[]*string `field:"optional" json:"outputModalities" yaml:"outputModalities"`
}

