package awslakeformation


// A structure for a data location object where permissions are granted or revoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &DataLocationResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html
//
type CfnPrincipalPermissions_DataLocationResourceProperty struct {
	// The identifier for the Data Catalog where the location is registered with AWS Lake Formation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html#cfn-lakeformation-principalpermissions-datalocationresource-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The Amazon Resource Name (ARN) that uniquely identifies the data location resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html#cfn-lakeformation-principalpermissions-datalocationresource-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

