package awsiottwinmaker


// The data connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataConnectorProperty := &DataConnectorProperty{
//   	IsNative: jsii.Boolean(false),
//   	Lambda: &LambdaFunctionProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-dataconnector.html
//
type CfnComponentType_DataConnectorProperty struct {
	// A boolean value that specifies whether the data connector is native to IoT TwinMaker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-dataconnector.html#cfn-iottwinmaker-componenttype-dataconnector-isnative
	//
	IsNative interface{} `field:"optional" json:"isNative" yaml:"isNative"`
	// The Lambda function associated with the data connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-dataconnector.html#cfn-iottwinmaker-componenttype-dataconnector-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

