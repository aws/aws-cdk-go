package interfacesawsmgn


// A reference to a Wave resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waveReference := &WaveReference{
//   	WaveArn: jsii.String("waveArn"),
//   }
//
type WaveReference struct {
	// The Arn of the Wave resource.
	WaveArn *string `field:"required" json:"waveArn" yaml:"waveArn"`
}

