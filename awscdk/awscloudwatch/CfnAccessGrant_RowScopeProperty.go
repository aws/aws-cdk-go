package awscloudwatch


// A single additive row-level filter on an access grant.
//
// A row is visible when its value for the given field matches the filter's values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowScopeProperty := &RowScopeProperty{
//   	Field: jsii.String("field"),
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-rowscope.html
//
type CfnAccessGrant_RowScopeProperty struct {
	// The field the allowlist applies to, for example serviceName or accountId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-rowscope.html#cfn-cloudwatch-accessgrant-rowscope-field
	//
	Field *string `field:"required" json:"field" yaml:"field"`
	// Match operator applied to this filter's values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-rowscope.html#cfn-cloudwatch-accessgrant-rowscope-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The values the field is matched against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-rowscope.html#cfn-cloudwatch-accessgrant-rowscope-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

