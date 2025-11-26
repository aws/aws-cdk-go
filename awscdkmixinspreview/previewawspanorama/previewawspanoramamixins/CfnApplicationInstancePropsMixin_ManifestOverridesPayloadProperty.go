package previewawspanoramamixins


// Parameter overrides for an application instance.
//
// This is a JSON document that has a single key ( `PayloadData` ) where the value is an escaped string representation of the overrides document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   manifestOverridesPayloadProperty := &ManifestOverridesPayloadProperty{
//   	PayloadData: jsii.String("payloadData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-applicationinstance-manifestoverridespayload.html
//
type CfnApplicationInstancePropsMixin_ManifestOverridesPayloadProperty struct {
	// The overrides document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-panorama-applicationinstance-manifestoverridespayload.html#cfn-panorama-applicationinstance-manifestoverridespayload-payloaddata
	//
	PayloadData *string `field:"optional" json:"payloadData" yaml:"payloadData"`
}

