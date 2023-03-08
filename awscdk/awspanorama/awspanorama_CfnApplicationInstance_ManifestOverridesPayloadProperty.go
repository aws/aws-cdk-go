package awspanorama


// Parameter overrides for an application instance.
//
// This is a JSON document that has a single key ( `PayloadData` ) where the value is an escaped string representation of the overrides document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestOverridesPayloadProperty := &manifestOverridesPayloadProperty{
//   	payloadData: jsii.String("payloadData"),
//   }
//
type CfnApplicationInstance_ManifestOverridesPayloadProperty struct {
	// The overrides document.
	PayloadData *string `field:"optional" json:"payloadData" yaml:"payloadData"`
}

