package interfacesawssso


// A reference to a InstanceAccessControlAttributeConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceAccessControlAttributeConfigurationReference := &InstanceAccessControlAttributeConfigurationReference{
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
type InstanceAccessControlAttributeConfigurationReference struct {
	// The InstanceArn of the InstanceAccessControlAttributeConfiguration resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
}

