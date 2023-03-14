package awswafregional


// Properties for defining a `CfnGeoMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGeoMatchSetProps := &cfnGeoMatchSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	geoMatchConstraints: []interface{}{
//   		&geoMatchConstraintProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGeoMatchSetProps struct {
	// A friendly name or description of the `GeoMatchSet` .
	//
	// You can't change the name of an `GeoMatchSet` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of `GeoMatchConstraint` objects, which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints interface{} `field:"optional" json:"geoMatchConstraints" yaml:"geoMatchConstraints"`
}

