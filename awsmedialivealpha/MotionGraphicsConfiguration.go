package awsmedialivealpha


// Motion graphics overlay configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var motionGraphicsInsertion MotionGraphicsInsertion
//
//   motionGraphicsConfiguration := &MotionGraphicsConfiguration{
//   	MotionGraphicsInsertion: motionGraphicsInsertion,
//   }
//
// Experimental.
type MotionGraphicsConfiguration struct {
	// Whether to enable the motion graphics overlay.
	// Default: MotionGraphicsInsertion.DISABLED
	//
	// Experimental.
	MotionGraphicsInsertion MotionGraphicsInsertion `field:"optional" json:"motionGraphicsInsertion" yaml:"motionGraphicsInsertion"`
}

