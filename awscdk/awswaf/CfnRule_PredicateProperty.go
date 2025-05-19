package awswaf


// Specifies the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , and `SizeConstraintSet` objects that you want to add to a `Rule` and, for each object, indicates whether you want to negate the settings, for example, requests that do NOT originate from the IP address 192.0.2.44.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predicateProperty := &PredicateProperty{
//   	DataId: jsii.String("dataId"),
//   	Negated: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicate.html
//
type CfnRule_PredicateProperty struct {
	// A unique identifier for a predicate in a `Rule` , such as `ByteMatchSetId` or `IPSetId` .
	//
	// The ID is returned by the corresponding `Create` or `List` command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicate.html#cfn-waf-rule-predicate-dataid
	//
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Set `Negated` to `False` if you want AWS WAF to allow, block, or count requests based on the settings in the specified `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` .
	//
	// For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow or block requests based on that IP address.
	//
	// Set `Negated` to `True` if you want AWS WAF to allow or block a request based on the negation of the settings in the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` . For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow, block, or count requests based on all IP addresses *except* `192.0.2.44` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicate.html#cfn-waf-rule-predicate-negated
	//
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// The type of predicate in a `Rule` , such as `ByteMatch` or `IPSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicate.html#cfn-waf-rule-predicate-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

