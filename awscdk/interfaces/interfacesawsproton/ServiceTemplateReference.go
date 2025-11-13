package interfacesawsproton


// A reference to a ServiceTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceTemplateReference := &ServiceTemplateReference{
//   	ServiceTemplateArn: jsii.String("serviceTemplateArn"),
//   }
//
type ServiceTemplateReference struct {
	// The Arn of the ServiceTemplate resource.
	ServiceTemplateArn *string `field:"required" json:"serviceTemplateArn" yaml:"serviceTemplateArn"`
}

