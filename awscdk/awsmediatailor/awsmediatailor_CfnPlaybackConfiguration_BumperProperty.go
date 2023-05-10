package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bumperProperty := &BumperProperty{
//   	EndUrl: jsii.String("endUrl"),
//   	StartUrl: jsii.String("startUrl"),
//   }
//
type CfnPlaybackConfiguration_BumperProperty struct {
	// `CfnPlaybackConfiguration.BumperProperty.EndUrl`.
	EndUrl *string `field:"optional" json:"endUrl" yaml:"endUrl"`
	// `CfnPlaybackConfiguration.BumperProperty.StartUrl`.
	StartUrl *string `field:"optional" json:"startUrl" yaml:"startUrl"`
}

