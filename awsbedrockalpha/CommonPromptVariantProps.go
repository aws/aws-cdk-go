package awsbedrockalpha


// Common properties for all prompt variants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var bedrockInvokable iBedrockInvokable
//
//   commonPromptVariantProps := &CommonPromptVariantProps{
//   	Model: bedrockInvokable,
//   	VariantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	PromptVariables: []*string{
//   		jsii.String("promptVariables"),
//   	},
//   }
//
// Experimental.
type CommonPromptVariantProps struct {
	// The model which is used to run the prompt.
	//
	// The model could be a foundation
	// model, a custom model, or a provisioned model.
	// Experimental.
	Model IBedrockInvokable `field:"required" json:"model" yaml:"model"`
	// The name of the prompt variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The variables in the prompt template that can be filled in at runtime.
	// Default: - No variables defined.
	//
	// Experimental.
	PromptVariables *[]*string `field:"optional" json:"promptVariables" yaml:"promptVariables"`
}

