package awsiottwinmaker


// The function body.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionProperty := &FunctionProperty{
//   	ImplementedBy: &DataConnectorProperty{
//   		IsNative: jsii.Boolean(false),
//   		Lambda: &LambdaFunctionProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	RequiredProperties: []*string{
//   		jsii.String("requiredProperties"),
//   	},
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-function.html
//
type CfnComponentType_FunctionProperty struct {
	// The data connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-function.html#cfn-iottwinmaker-componenttype-function-implementedby
	//
	ImplementedBy interface{} `field:"optional" json:"implementedBy" yaml:"implementedBy"`
	// The required properties of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-function.html#cfn-iottwinmaker-componenttype-function-requiredproperties
	//
	RequiredProperties *[]*string `field:"optional" json:"requiredProperties" yaml:"requiredProperties"`
	// The scope of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-function.html#cfn-iottwinmaker-componenttype-function-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

