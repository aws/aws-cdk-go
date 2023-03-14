package awsiottwinmaker


// The function body.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionProperty := &functionProperty{
//   	implementedBy: &dataConnectorProperty{
//   		isNative: jsii.Boolean(false),
//   		lambda: &lambdaFunctionProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	requiredProperties: []*string{
//   		jsii.String("requiredProperties"),
//   	},
//   	scope: jsii.String("scope"),
//   }
//
type CfnComponentType_FunctionProperty struct {
	// The data connector.
	ImplementedBy interface{} `field:"optional" json:"implementedBy" yaml:"implementedBy"`
	// The required properties of the function.
	RequiredProperties *[]*string `field:"optional" json:"requiredProperties" yaml:"requiredProperties"`
	// The scope of the function.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

