package interfacesawsdynamodb


// A reference to a GlobalTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalTableReference := &GlobalTableReference{
//   	GlobalTableArn: jsii.String("globalTableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
type GlobalTableReference struct {
	// The ARN of the GlobalTable resource.
	GlobalTableArn *string `field:"required" json:"globalTableArn" yaml:"globalTableArn"`
	// The TableName of the GlobalTable resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

