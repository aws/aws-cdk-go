package mixinsawsappmesh


// An object that represents the key value pairs for the JSON.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jsonFormatRefProperty := &JsonFormatRefProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-jsonformatref.html
//
type CfnVirtualNodePropsMixin_JsonFormatRefProperty struct {
	// The specified key for the JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-jsonformatref.html#cfn-appmesh-virtualnode-jsonformatref-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The specified value for the JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-jsonformatref.html#cfn-appmesh-virtualnode-jsonformatref-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

