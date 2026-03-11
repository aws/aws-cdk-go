package awslakeformation


// The AWS Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataLakePrincipalProperty := &DataLakePrincipalProperty{
//   	DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalakeprincipal.html
//
type CfnPrincipalPermissionsPropsMixin_DataLakePrincipalProperty struct {
	// An identifier for the AWS Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalakeprincipal.html#cfn-lakeformation-principalpermissions-datalakeprincipal-datalakeprincipalidentifier
	//
	DataLakePrincipalIdentifier interface{} `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

