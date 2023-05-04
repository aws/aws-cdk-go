package awslakeformation


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
type CfnDataLakeSettings_PrincipalPermissionsProperty struct {
	// `CfnDataLakeSettings.PrincipalPermissionsProperty.Permissions`.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// `CfnDataLakeSettings.PrincipalPermissionsProperty.Principal`.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

