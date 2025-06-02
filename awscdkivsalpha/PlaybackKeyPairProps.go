package awscdkivsalpha


// Properties for creating a new Playback Key Pair.
//
// Example:
//   keyPair := ivs.NewPlaybackKeyPair(this, jsii.String("PlaybackKeyPair"), &PlaybackKeyPairProps{
//   	PublicKeyMaterial: myPublicKeyPemString,
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
	// Default: Automatically generated name.
	//
	// Experimental.
	PlaybackKeyPairName *string `field:"optional" json:"playbackKeyPairName" yaml:"playbackKeyPairName"`
}

