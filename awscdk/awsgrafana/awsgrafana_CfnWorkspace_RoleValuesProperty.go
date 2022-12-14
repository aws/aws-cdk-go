package awsgrafana


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roleValuesProperty := &roleValuesProperty{
//   	admin: []*string{
//   		jsii.String("admin"),
//   	},
//   	editor: []*string{
//   		jsii.String("editor"),
//   	},
//   }
//
type CfnWorkspace_RoleValuesProperty struct {
	// `CfnWorkspace.RoleValuesProperty.Admin`.
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// `CfnWorkspace.RoleValuesProperty.Editor`.
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

