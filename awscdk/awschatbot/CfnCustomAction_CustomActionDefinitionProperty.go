package awschatbot


// The definition of the command to run when invoked as an alias or as an action button.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionDefinitionProperty := &CustomActionDefinitionProperty{
//   	CommandText: jsii.String("commandText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html
//
type CfnCustomAction_CustomActionDefinitionProperty struct {
	// The command string to run which may include variables by prefixing with a dollar sign ($).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html#cfn-chatbot-customaction-customactiondefinition-commandtext
	//
	CommandText *string `field:"required" json:"commandText" yaml:"commandText"`
}

