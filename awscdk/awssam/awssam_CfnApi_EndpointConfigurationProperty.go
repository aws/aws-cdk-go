package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	type: jsii.String("type"),
//   	vpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
type CfnApi_EndpointConfigurationProperty struct {
	// `CfnApi.EndpointConfigurationProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// `CfnApi.EndpointConfigurationProperty.VpcEndpointIds`.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

