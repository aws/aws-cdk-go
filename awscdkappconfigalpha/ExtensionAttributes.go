package awscdkappconfigalpha


// Attributes of an existing AWS AppConfig extension to import.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//
//   var action action
//
//   extensionAttributes := &ExtensionAttributes{
//   	ExtensionId: jsii.String("extensionId"),
//   	ExtensionVersionNumber: jsii.Number(123),
//
//   	// the properties below are optional
//   	Actions: []*action{
//   		action,
//   	},
//   	Description: jsii.String("description"),
//   	ExtensionArn: jsii.String("extensionArn"),
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type ExtensionAttributes struct {
	// The ID of the extension.
	// Experimental.
	ExtensionId *string `field:"required" json:"extensionId" yaml:"extensionId"`
	// The version number of the extension.
	// Experimental.
	ExtensionVersionNumber *float64 `field:"required" json:"extensionVersionNumber" yaml:"extensionVersionNumber"`
	// The actions of the extension.
	// Experimental.
	Actions *[]Action `field:"optional" json:"actions" yaml:"actions"`
	// The description of the extension.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the extension.
	// Experimental.
	ExtensionArn *string `field:"optional" json:"extensionArn" yaml:"extensionArn"`
	// The name of the extension.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

