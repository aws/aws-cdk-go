package awsroute53


// Properties for defining a `CfnDNSSEC`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDNSSECProps := &cfnDNSSECProps{
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type CfnDNSSECProps struct {
	// A unique string (ID) that is used to identify a hosted zone.
	//
	// For example: `Z00001111A1ABCaaABC11` .
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

