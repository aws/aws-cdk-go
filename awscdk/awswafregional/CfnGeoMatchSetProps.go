package awswafregional


// Properties for defining a `CfnGeoMatchSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGeoMatchSetProps := &CfnGeoMatchSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	GeoMatchConstraints: []interface{}{
//   		&GeoMatchConstraintProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-geomatchset.html
//
type CfnGeoMatchSetProps struct {
	// A friendly name or description of the `GeoMatchSet` .
	//
	// You can't change the name of an `GeoMatchSet` after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-geomatchset.html#cfn-wafregional-geomatchset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of `GeoMatchConstraint` objects, which contain the country that you want AWS WAF to search for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-geomatchset.html#cfn-wafregional-geomatchset-geomatchconstraints
	//
	GeoMatchConstraints interface{} `field:"optional" json:"geoMatchConstraints" yaml:"geoMatchConstraints"`
}

