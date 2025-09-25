package awsconnect


// A reference to a ContactFlowModule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactFlowModuleReference := &ContactFlowModuleReference{
//   	ContactFlowModuleArn: jsii.String("contactFlowModuleArn"),
//   }
//
type ContactFlowModuleReference struct {
	// The ContactFlowModuleArn of the ContactFlowModule resource.
	ContactFlowModuleArn *string `field:"required" json:"contactFlowModuleArn" yaml:"contactFlowModuleArn"`
}

