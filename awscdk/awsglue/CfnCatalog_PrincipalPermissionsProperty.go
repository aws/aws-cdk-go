package awsglue


// Permissions granted to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalPermissionsProperty := &PrincipalPermissionsProperty{
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	Principal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-principalpermissions.html
//
type CfnCatalog_PrincipalPermissionsProperty struct {
	// The permissions that are granted to the principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-principalpermissions.html#cfn-glue-catalog-principalpermissions-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-principalpermissions.html#cfn-glue-catalog-principalpermissions-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

