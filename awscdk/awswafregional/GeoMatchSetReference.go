package awswafregional


// A reference to a GeoMatchSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoMatchSetReference := &GeoMatchSetReference{
//   	GeoMatchSetId: jsii.String("geoMatchSetId"),
//   }
//
type GeoMatchSetReference struct {
	// The Id of the GeoMatchSet resource.
	GeoMatchSetId *string `field:"required" json:"geoMatchSetId" yaml:"geoMatchSetId"`
}

