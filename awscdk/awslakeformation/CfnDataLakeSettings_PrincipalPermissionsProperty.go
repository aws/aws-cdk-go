package awslakeformation


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-principalpermissions.html
//
type CfnDataLakeSettings_PrincipalPermissionsProperty struct {
	// The permissions that are granted to the principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-principalpermissions.html#cfn-lakeformation-datalakesettings-principalpermissions-permissions
	//
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// The principal who is granted permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-principalpermissions.html#cfn-lakeformation-datalakesettings-principalpermissions-principal
	//
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
}

