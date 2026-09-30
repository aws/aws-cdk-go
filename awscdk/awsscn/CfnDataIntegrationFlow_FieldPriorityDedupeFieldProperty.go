package awsscn


// A deduplication field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldPriorityDedupeFieldProperty := &FieldPriorityDedupeFieldProperty{
//   	Name: jsii.String("name"),
//   	SortOrder: jsii.String("sortOrder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-fieldprioritydedupefield.html
//
type CfnDataIntegrationFlow_FieldPriorityDedupeFieldProperty struct {
	// The name of the deduplication field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-fieldprioritydedupefield.html#cfn-scn-dataintegrationflow-fieldprioritydedupefield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The sort order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-fieldprioritydedupefield.html#cfn-scn-dataintegrationflow-fieldprioritydedupefield-sortorder
	//
	SortOrder *string `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

