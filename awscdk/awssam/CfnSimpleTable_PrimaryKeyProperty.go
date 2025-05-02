package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryKeyProperty := &PrimaryKeyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-primarykey.html
//
type CfnSimpleTable_PrimaryKeyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-primarykey.html#cfn-serverless-simpletable-primarykey-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-simpletable-primarykey.html#cfn-serverless-simpletable-primarykey-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

