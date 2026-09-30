package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vectorAttributeProperty := &VectorAttributeProperty{
//   	AttributeName: jsii.String("attributeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorattribute.html
//
type CfnGlobalTablePropsMixin_VectorAttributeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorattribute.html#cfn-dynamodb-globaltable-vectorattribute-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
}

