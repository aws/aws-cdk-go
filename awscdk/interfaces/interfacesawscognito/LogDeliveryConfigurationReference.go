package interfacesawscognito


// A reference to a LogDeliveryConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryConfigurationReference := &LogDeliveryConfigurationReference{
//   	LogDeliveryConfigurationId: jsii.String("logDeliveryConfigurationId"),
//   }
//
type LogDeliveryConfigurationReference struct {
	// The Id of the LogDeliveryConfiguration resource.
	LogDeliveryConfigurationId *string `field:"required" json:"logDeliveryConfigurationId" yaml:"logDeliveryConfigurationId"`
}

