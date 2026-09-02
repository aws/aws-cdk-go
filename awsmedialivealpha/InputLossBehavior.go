package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Behavior on input loss: substitute black, optionally repeat the last frame, then show a solid color or a slate image.
//
// Example:
//   var slateBucket IBucket
//
//
//   inputLoss := &InputLossBehavior{
//   	BlackFrame: awscdk.Duration_Seconds(jsii.Number(1)),
//   	RepeatFrame: awscdk.Duration_*Seconds(jsii.Number(5)),
//   	ImageType: medialive.InputLossImageType_SLATE(),
//   	ImageSlate: medialive.FileLocation_FromBucket(slateBucket, jsii.String("slates/offline.png")),
//   }
//
// Experimental.
type InputLossBehavior struct {
	// How long to substitute black before showing the input-loss image (up to Duration.seconds(1000); Duration.seconds(1000) is interpreted as infinite).
	// Default: - service default.
	//
	// Experimental.
	BlackFrame awscdk.Duration `field:"optional" json:"blackFrame" yaml:"blackFrame"`
	// The image color as 6 hex characters (RGB).
	//
	// Used when InputLossImageType.COLOR.
	// Default: - service default.
	//
	// Experimental.
	ImageColor *string `field:"optional" json:"imageColor" yaml:"imageColor"`
	// The slate image to display.
	//
	// Used when `imageType` is SLATE. Provide a `FileLocation`
	// referencing an S3 bucket (`FileLocation.fromBucket`, which auto-grants read access) or
	// a URL (`FileLocation.url`).
	// Default: - service default.
	//
	// Experimental.
	ImageSlate FileLocation `field:"optional" json:"imageSlate" yaml:"imageSlate"`
	// Whether to substitute a solid color or a slate image after the black period.
	// Default: - service default.
	//
	// Experimental.
	ImageType InputLossImageType `field:"optional" json:"imageType" yaml:"imageType"`
	// How long to repeat the previous picture before substituting black (up to Duration.seconds(1000); Duration.seconds(1000) is interpreted as infinite).
	// Default: - service default.
	//
	// Experimental.
	RepeatFrame awscdk.Duration `field:"optional" json:"repeatFrame" yaml:"repeatFrame"`
}

