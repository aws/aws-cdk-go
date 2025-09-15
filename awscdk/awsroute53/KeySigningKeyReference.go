package awsroute53


// A reference to a KeySigningKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keySigningKeyReference := &KeySigningKeyReference{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	KeySigningKeyName: jsii.String("keySigningKeyName"),
//   }
//
type KeySigningKeyReference struct {
	// The HostedZoneId of the KeySigningKey resource.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The Name of the KeySigningKey resource.
	KeySigningKeyName *string `field:"required" json:"keySigningKeyName" yaml:"keySigningKeyName"`
}

