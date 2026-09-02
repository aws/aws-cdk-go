package interfacesawsssm


// A reference to a CloudConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudConnectorReference := &CloudConnectorReference{
//   	CloudConnectorArn: jsii.String("cloudConnectorArn"),
//   }
//
type CloudConnectorReference struct {
	// The CloudConnectorArn of the CloudConnector resource.
	CloudConnectorArn *string `field:"required" json:"cloudConnectorArn" yaml:"cloudConnectorArn"`
}

