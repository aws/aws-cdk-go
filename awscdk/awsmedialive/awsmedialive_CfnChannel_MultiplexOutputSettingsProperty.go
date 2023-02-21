package awsmedialive


// Configuration of a Multiplex output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexOutputSettingsProperty := &multiplexOutputSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
type CfnChannel_MultiplexOutputSettingsProperty struct {
	// Destination is a Multiplex.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

