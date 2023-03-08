// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   tableAttributes := &TableAttributes{
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// Experimental.
type TableAttributes struct {
	// Experimental.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

