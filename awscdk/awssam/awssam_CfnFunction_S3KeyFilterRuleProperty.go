package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterRuleProperty := &s3KeyFilterRuleProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnFunction_S3KeyFilterRuleProperty struct {
	// `CfnFunction.S3KeyFilterRuleProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnFunction.S3KeyFilterRuleProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

