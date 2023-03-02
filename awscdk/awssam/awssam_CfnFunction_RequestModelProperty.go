package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestModelProperty := &requestModelProperty{
//   	model: jsii.String("model"),
//
//   	// the properties below are optional
//   	required: jsii.Boolean(false),
//   	validateBody: jsii.Boolean(false),
//   	validateParameters: jsii.Boolean(false),
//   }
//
type CfnFunction_RequestModelProperty struct {
	// `CfnFunction.RequestModelProperty.Model`.
	Model *string `field:"required" json:"model" yaml:"model"`
	// `CfnFunction.RequestModelProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// `CfnFunction.RequestModelProperty.ValidateBody`.
	ValidateBody interface{} `field:"optional" json:"validateBody" yaml:"validateBody"`
	// `CfnFunction.RequestModelProperty.ValidateParameters`.
	ValidateParameters interface{} `field:"optional" json:"validateParameters" yaml:"validateParameters"`
}

