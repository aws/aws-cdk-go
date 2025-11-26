package previewawscassandramixins


// Properties for CfnTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTypeMixinProps := &CfnTypeMixinProps{
//   	Fields: []interface{}{
//   		&FieldProperty{
//   			FieldName: jsii.String("fieldName"),
//   			FieldType: jsii.String("fieldType"),
//   		},
//   	},
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	TypeName: jsii.String("typeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-type.html
//
type CfnTypeMixinProps struct {
	// A list of fields that define this type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-type.html#cfn-cassandra-type-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the keyspace to create the type in.
	//
	// The keyspace must already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-type.html#cfn-cassandra-type-keyspacename
	//
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// The name of the user-defined type.
	//
	// UDT names must contain 48 characters or less, must begin with an alphabetic character, and can only contain alpha-numeric characters and underscores. Amazon Keyspaces converts upper case characters automatically into lower case characters. For more information, see [Create a user-defined type (UDT) in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/keyspaces-create-udt.html) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-type.html#cfn-cassandra-type-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

