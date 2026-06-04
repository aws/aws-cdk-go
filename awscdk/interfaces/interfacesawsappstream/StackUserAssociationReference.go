package interfacesawsappstream


// A reference to a StackUserAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackUserAssociationReference := &StackUserAssociationReference{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	StackName: jsii.String("stackName"),
//   	UserName: jsii.String("userName"),
//   }
//
type StackUserAssociationReference struct {
	// The AuthenticationType of the StackUserAssociation resource.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The StackName of the StackUserAssociation resource.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The UserName of the StackUserAssociation resource.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
}

