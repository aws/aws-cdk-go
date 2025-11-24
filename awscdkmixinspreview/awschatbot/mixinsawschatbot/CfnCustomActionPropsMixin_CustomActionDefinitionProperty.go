package mixinsawschatbot


// > AWS Chatbot is now  . [Learn more](https://docs.aws.amazon.com//chatbot/latest/adminguide/service-rename.html) >  > `Type` attribute values remain unchanged.
//
// The definition of the command to run when invoked as an alias or as an action button.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionDefinitionProperty := &CustomActionDefinitionProperty{
//   	CommandText: jsii.String("commandText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html
//
type CfnCustomActionPropsMixin_CustomActionDefinitionProperty struct {
	// The command string to run which may include variables by prefixing with a dollar sign ($).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html#cfn-chatbot-customaction-customactiondefinition-commandtext
	//
	CommandText *string `field:"optional" json:"commandText" yaml:"commandText"`
}

