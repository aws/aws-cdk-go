package awsimagebuilderalpha


// The definition of the parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   var default_ interface{}
//
//   componentDocumentParameterDefinition := &ComponentDocumentParameterDefinition{
//   	Type: imagebuilder_alpha.ComponentParameterType_STRING,
//
//   	// the properties below are optional
//   	Default: default_,
//   	Description: jsii.String("description"),
//   }
//
// Experimental.
type ComponentDocumentParameterDefinition struct {
	// The type of the parameter.
	// Experimental.
	Type ComponentParameterType `field:"required" json:"type" yaml:"type"`
	// The default value of the parameter.
	// Default: - none, indicating the parameter is required.
	//
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// The description of the parameter.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

