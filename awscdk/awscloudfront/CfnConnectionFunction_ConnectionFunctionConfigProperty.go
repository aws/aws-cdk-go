package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionFunctionConfigProperty := &ConnectionFunctionConfigProperty{
//   	Comment: jsii.String("comment"),
//   	Runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	KeyValueStoreAssociations: []interface{}{
//   		&KeyValueStoreAssociationProperty{
//   			KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html
//
type CfnConnectionFunction_ConnectionFunctionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-comment
	//
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-keyvaluestoreassociations
	//
	KeyValueStoreAssociations interface{} `field:"optional" json:"keyValueStoreAssociations" yaml:"keyValueStoreAssociations"`
}

