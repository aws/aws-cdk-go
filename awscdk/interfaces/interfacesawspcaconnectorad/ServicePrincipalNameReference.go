package interfacesawspcaconnectorad


// A reference to a ServicePrincipalName resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   servicePrincipalNameReference := &ServicePrincipalNameReference{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	DirectoryRegistrationArn: jsii.String("directoryRegistrationArn"),
//   }
//
type ServicePrincipalNameReference struct {
	// The ConnectorArn of the ServicePrincipalName resource.
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
	// The DirectoryRegistrationArn of the ServicePrincipalName resource.
	DirectoryRegistrationArn *string `field:"required" json:"directoryRegistrationArn" yaml:"directoryRegistrationArn"`
}

