package awsroute53


// The status for a Key Signing Key.
//
// Example:
//   var hostedZone hostedZone
//   var kmsKey key
//
//   route53.NewKeySigningKey(this, jsii.String("KeySigningKey"), &KeySigningKeyProps{
//   	HostedZone: HostedZone,
//   	KmsKey: KmsKey,
//   	KeySigningKeyName: jsii.String("ksk"),
//   	Status: route53.KeySigningKeyStatus_ACTIVE,
//   })
//
type KeySigningKeyStatus string

const (
	// The KSK is being used for signing.
	KeySigningKeyStatus_ACTIVE KeySigningKeyStatus = "ACTIVE"
	// The KSK is not being used for signing.
	KeySigningKeyStatus_INACTIVE KeySigningKeyStatus = "INACTIVE"
)

