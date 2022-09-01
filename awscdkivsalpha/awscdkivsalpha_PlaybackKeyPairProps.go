// The CDK Construct Library for AWS::IVS
package awscdkivsalpha


// Properties for creating a new Playback Key Pair.
//
// Example:
//   keyPair := ivs.NewPlaybackKeyPair(this, jsii.String("PlaybackKeyPair"), &playbackKeyPairProps{
//   	publicKeyMaterial: myPublicKeyPemString,
//   })
//
// Experimental.
type PlaybackKeyPairProps struct {
	// The public portion of a customer-generated key pair.
	// Experimental.
	PublicKeyMaterial *string `field:"required" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource.
	//
	// The value does not need to be unique.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

