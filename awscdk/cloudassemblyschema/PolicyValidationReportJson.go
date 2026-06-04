package cloudassemblyschema


// The top-level structure of the policy validation report file.
//
// Example:
//   import "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//
//   report := &PolicyValidationReportJson{
//   	Version: jsii.String("1.0"),
//   	PluginReports: []PluginReportJson{
//   		&PluginReportJson{
//   			PluginName: jsii.String("my-plugin"),
//   			Conclusion: jsii.String("success"),
//   			Violations: []PolicyViolationJson{
//   			},
//   		},
//   	},
//   }
//
type PolicyValidationReportJson struct {
	// Reports from all validation plugins that ran during synthesis.
	PluginReports *[]*PluginReportJson `field:"required" json:"pluginReports" yaml:"pluginReports"`
	// Protocol version.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Report title, if present.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

