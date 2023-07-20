package awsappmesh


// An object that represents the key value pairs for the JSON.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFormatRefProperty := &JsonFormatRefProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html
//
type CfnVirtualGateway_JsonFormatRefProperty struct {
	// The specified key for the JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html#cfn-appmesh-virtualgateway-jsonformatref-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The specified value for the JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-jsonformatref.html#cfn-appmesh-virtualgateway-jsonformatref-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

