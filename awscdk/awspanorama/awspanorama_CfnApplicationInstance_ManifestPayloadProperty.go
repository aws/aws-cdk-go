package awspanorama


// A application verion's manifest file.
//
// This is a JSON document that has a single key ( `PayloadData` ) where the value is an escaped string representation of the application manifest ( `graph.json` ). This file is located in the `graphs` folder in your application source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestPayloadProperty := &manifestPayloadProperty{
//   	payloadData: jsii.String("payloadData"),
//   }
//
type CfnApplicationInstance_ManifestPayloadProperty struct {
	// The application manifest.
	PayloadData *string `field:"optional" json:"payloadData" yaml:"payloadData"`
}

