package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableAttributes := &tableAttributes{
//   	tableArn: jsii.String("tableArn"),
//   	tableName: jsii.String("tableName"),
//   }
//
// Experimental.
type TableAttributes struct {
	// Experimental.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

