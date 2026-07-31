package awsmediaconnectalpha


// Properties for MediaLive Channel Router Input configuration without a specific channel connection.
//
// Use this when you want to set up the router input before the target MediaLive channel exists.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var secret Secret
//
//   mediaLiveChannelConfigurationWithoutConnectionProps := &MediaLiveChannelConfigurationWithoutConnectionProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//
//   	// the properties below are optional
//   	SourceTransitDecryption: &TransitEncryption{
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   }
//
// Experimental.
type MediaLiveChannelConfigurationWithoutConnectionProps struct {
	// Availability Zone the router input will be placed in.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

