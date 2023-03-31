package awsroute53resolver


// Properties for defining a `CfnResolverDNSSECConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverDNSSECConfigProps := &cfnResolverDNSSECConfigProps{
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverDNSSECConfigProps struct {
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

