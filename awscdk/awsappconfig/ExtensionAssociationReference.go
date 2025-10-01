package awsappconfig


// A reference to a ExtensionAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionAssociationReference := &ExtensionAssociationReference{
//   	ExtensionAssociationArn: jsii.String("extensionAssociationArn"),
//   	ExtensionAssociationId: jsii.String("extensionAssociationId"),
//   }
//
type ExtensionAssociationReference struct {
	// The ARN of the ExtensionAssociation resource.
	ExtensionAssociationArn *string `field:"required" json:"extensionAssociationArn" yaml:"extensionAssociationArn"`
	// The Id of the ExtensionAssociation resource.
	ExtensionAssociationId *string `field:"required" json:"extensionAssociationId" yaml:"extensionAssociationId"`
}

