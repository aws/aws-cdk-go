package awsmedialivealpha


// Properties for a multicast input.
//
// Requires `anywhereSettings` on the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var multicastProtocol MulticastProtocol
//
//   multicastInputProps := &MulticastInputProps{
//   	Sources: []MulticastInputSource{
//   		&MulticastInputSource{
//   			Address: jsii.String("address"),
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			Protocol: multicastProtocol,
//   			SourceIp: jsii.String("sourceIp"),
//   		},
//   	},
//   }
//
// Experimental.
type MulticastInputProps struct {
	// The multicast sources.
	//
	// Provide two for a STANDARD (redundant-pipeline) channel.
	// Experimental.
	Sources *[]*MulticastInputSource `field:"required" json:"sources" yaml:"sources"`
}

