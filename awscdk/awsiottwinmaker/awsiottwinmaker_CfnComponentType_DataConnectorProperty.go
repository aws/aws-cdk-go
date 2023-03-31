package awsiottwinmaker


// The data connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataConnectorProperty := &dataConnectorProperty{
//   	isNative: jsii.Boolean(false),
//   	lambda: &lambdaFunctionProperty{
//   		arn: jsii.String("arn"),
//   	},
//   }
//
type CfnComponentType_DataConnectorProperty struct {
	// A boolean value that specifies whether the data connector is native to IoT TwinMaker.
	IsNative interface{} `field:"optional" json:"isNative" yaml:"isNative"`
	// The Lambda function associated with the data connector.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

