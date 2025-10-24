package awsroute53


// The attributes of a key signing key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone HostedZone
//
//   keySigningKeyAttributes := &KeySigningKeyAttributes{
//   	HostedZone: hostedZone,
//   	KeySigningKeyName: jsii.String("keySigningKeyName"),
//   }
//
type KeySigningKeyAttributes struct {
	// The hosted zone that the key signing key signs.
	HostedZone IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// The name of the key signing key.
	KeySigningKeyName *string `field:"required" json:"keySigningKeyName" yaml:"keySigningKeyName"`
}

