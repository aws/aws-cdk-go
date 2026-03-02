package interfacesawsconnect


// A reference to a ContactFlowModuleAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactFlowModuleAliasReference := &ContactFlowModuleAliasReference{
//   	ContactFlowModuleAliasArn: jsii.String("contactFlowModuleAliasArn"),
//   }
//
type ContactFlowModuleAliasReference struct {
	// The ContactFlowModuleAliasARN of the ContactFlowModuleAlias resource.
	ContactFlowModuleAliasArn *string `field:"required" json:"contactFlowModuleAliasArn" yaml:"contactFlowModuleAliasArn"`
}

