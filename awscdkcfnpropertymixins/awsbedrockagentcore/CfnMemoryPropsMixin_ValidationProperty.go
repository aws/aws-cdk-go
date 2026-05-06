package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   validationProperty := &ValidationProperty{
//   	NumberValidation: &NumberValidationProperty{
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//   	},
//   	StringListValidation: &StringListValidationProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		MaxItems: jsii.Number(123),
//   	},
//   	StringValidation: &StringValidationProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-validation.html
//
type CfnMemoryPropsMixin_ValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-validation.html#cfn-bedrockagentcore-memory-validation-numbervalidation
	//
	NumberValidation interface{} `field:"optional" json:"numberValidation" yaml:"numberValidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-validation.html#cfn-bedrockagentcore-memory-validation-stringlistvalidation
	//
	StringListValidation interface{} `field:"optional" json:"stringListValidation" yaml:"stringListValidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-validation.html#cfn-bedrockagentcore-memory-validation-stringvalidation
	//
	StringValidation interface{} `field:"optional" json:"stringValidation" yaml:"stringValidation"`
}

