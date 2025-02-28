package awscassandra


// The name and data type of an individual field in a user-defined type (UDT).
//
// In addition to a Cassandra data type, you can also use another UDT. When you nest another UDT or collection data type, you have to declare them with the `FROZEN` keyword.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldProperty := &FieldProperty{
//   	FieldName: jsii.String("fieldName"),
//   	FieldType: jsii.String("fieldType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-type-field.html
//
type CfnType_FieldProperty struct {
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-type-field.html#cfn-cassandra-type-field-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The data type of the field.
	//
	// This can be any Cassandra data type or another user-defined type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-type-field.html#cfn-cassandra-type-field-fieldtype
	//
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
}

