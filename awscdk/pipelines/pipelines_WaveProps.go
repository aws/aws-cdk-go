package pipelines


// Construction properties for a `Wave`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var step step
//
//   waveProps := &waveProps{
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   }
//
type WaveProps struct {
	// Additional steps to run after all of the stages in the wave.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

