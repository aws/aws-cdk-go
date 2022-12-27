package awsrds


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	port: jsii.String("port"),
//   }
//
type CfnDBInstance_EndpointProperty struct {
	// `CfnDBInstance.EndpointProperty.Address`.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// `CfnDBInstance.EndpointProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnDBInstance.EndpointProperty.Port`.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

