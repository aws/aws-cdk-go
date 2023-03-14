package awswafregional


// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// The country from which web requests originate that you want AWS WAF to search for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoMatchConstraintProperty := &geoMatchConstraintProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnGeoMatchSet_GeoMatchConstraintProperty struct {
	// The type of geographical area you want AWS WAF to search for.
	//
	// Currently `Country` is the only valid value.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The country that you want AWS WAF to search for.
	Value *string `field:"required" json:"value" yaml:"value"`
}

