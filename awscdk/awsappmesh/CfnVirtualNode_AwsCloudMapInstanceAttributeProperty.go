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
//   awsCloudMapInstanceAttributeProperty := &AwsCloudMapInstanceAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html
//
type CfnVirtualNode_AwsCloudMapInstanceAttributeProperty struct {
	// The name of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html#cfn-appmesh-virtualnode-awscloudmapinstanceattribute-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of an AWS Cloud Map service instance attribute key.
	//
	// Any AWS Cloud Map service instance that contains the specified key and value is returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.html#cfn-appmesh-virtualnode-awscloudmapinstanceattribute-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

