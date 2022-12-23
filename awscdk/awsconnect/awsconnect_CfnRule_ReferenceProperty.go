package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceProperty := &referenceProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnRule_ReferenceProperty struct {
	// `CfnRule.ReferenceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnRule.ReferenceProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

