package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueStoreAssociationProperty := &KeyValueStoreAssociationProperty{
//   	KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-keyvaluestoreassociation.html
//
type CfnConnectionFunction_KeyValueStoreAssociationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-keyvaluestoreassociation.html#cfn-cloudfront-connectionfunction-keyvaluestoreassociation-keyvaluestorearn
	//
	KeyValueStoreArn *string `field:"required" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
}

