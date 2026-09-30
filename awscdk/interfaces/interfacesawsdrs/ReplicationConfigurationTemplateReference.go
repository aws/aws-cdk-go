package interfacesawsdrs


// A reference to a ReplicationConfigurationTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationTemplateReference := &ReplicationConfigurationTemplateReference{
//   	ReplicationConfigurationTemplateArn: jsii.String("replicationConfigurationTemplateArn"),
//   }
//
type ReplicationConfigurationTemplateReference struct {
	// The Arn of the ReplicationConfigurationTemplate resource.
	ReplicationConfigurationTemplateArn *string `field:"required" json:"replicationConfigurationTemplateArn" yaml:"replicationConfigurationTemplateArn"`
}

