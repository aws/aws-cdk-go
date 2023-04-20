package awspinpoint


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDeliveryConfigurationProperty := &CustomDeliveryConfigurationProperty{
//   	DeliveryUri: jsii.String("deliveryUri"),
//   	EndpointTypes: []*string{
//   		jsii.String("endpointTypes"),
//   	},
//   }
//
type CfnCampaign_CustomDeliveryConfigurationProperty struct {
	// `CfnCampaign.CustomDeliveryConfigurationProperty.DeliveryUri`.
	DeliveryUri *string `field:"optional" json:"deliveryUri" yaml:"deliveryUri"`
	// `CfnCampaign.CustomDeliveryConfigurationProperty.EndpointTypes`.
	EndpointTypes *[]*string `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
}

