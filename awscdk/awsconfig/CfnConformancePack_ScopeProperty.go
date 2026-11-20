package awsconfig


// The global tag scope of the managed conformance pack.
//
// If specified, only resources with this tag key and value pair are in scope of the conformance pack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeProperty := &ScopeProperty{
//   	TagKey: jsii.String("tagKey"),
//
//   	// the properties below are optional
//   	TagValue: jsii.String("tagValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-scope.html
//
type CfnConformancePack_ScopeProperty struct {
	// The tag key of the tag scope.
	//
	// Only resources with this tag key are in scope of the conformance pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-scope.html#cfn-config-conformancepack-scope-tagkey
	//
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// The tag value of the tag scope.
	//
	// If omitted, resources are in scope based on the presence of the tag key alone, regardless of its value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-scope.html#cfn-config-conformancepack-scope-tagvalue
	//
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

