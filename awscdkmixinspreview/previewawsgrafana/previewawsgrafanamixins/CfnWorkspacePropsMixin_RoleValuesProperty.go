package previewawsgrafanamixins


// This structure defines which groups defined in the SAML assertion attribute are to be mapped to the Grafana `Admin` and `Editor` roles in the workspace.
//
// SAML authenticated users not part of `Admin` or `Editor` role groups have `Viewer` permission over the workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   roleValuesProperty := &RoleValuesProperty{
//   	Admin: []*string{
//   		jsii.String("admin"),
//   	},
//   	Editor: []*string{
//   		jsii.String("editor"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-rolevalues.html
//
type CfnWorkspacePropsMixin_RoleValuesProperty struct {
	// A list of groups from the SAML assertion attribute to grant the Grafana `Admin` role to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-rolevalues.html#cfn-grafana-workspace-rolevalues-admin
	//
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// A list of groups from the SAML assertion attribute to grant the Grafana `Editor` role to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-rolevalues.html#cfn-grafana-workspace-rolevalues-editor
	//
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

