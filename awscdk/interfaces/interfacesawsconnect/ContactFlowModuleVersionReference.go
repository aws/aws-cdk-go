package interfacesawsconnect


// A reference to a ContactFlowModuleVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactFlowModuleVersionReference := &ContactFlowModuleVersionReference{
//   	ContactFlowModuleVersionArn: jsii.String("contactFlowModuleVersionArn"),
//   }
//
type ContactFlowModuleVersionReference struct {
	// The ContactFlowModuleVersionARN of the ContactFlowModuleVersion resource.
	ContactFlowModuleVersionArn *string `field:"required" json:"contactFlowModuleVersionArn" yaml:"contactFlowModuleVersionArn"`
}

