package awsappmesh


// An object that represents the AWS Cloud Map attribute information for your virtual node.
//
// > AWS Cloud Map is not available in the eu-south-1 Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsCloudMapInstanceAttributeProperty := &awsCloudMapInstanceAttributeProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnVirtualNode_AwsCloudMapInstanceAttributeProperty struct {
	// The name of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	Value *string `field:"required" json:"value" yaml:"value"`
}

