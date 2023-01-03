package awsgrafana


// This structure defines which groups defined in the SAML assertion attribute are to be mapped to the Grafana `Admin` and `Editor` roles in the workspace.
//
// SAML authenticated users not part of `Admin` or `Editor` role groups have `Viewer` permission over the workspace.
//
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
	// A list of groups from the SAML assertion attribute to grant the Grafana `Admin` role to.
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// A list of groups from the SAML assertion attribute to grant the Grafana `Editor` role to.
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

