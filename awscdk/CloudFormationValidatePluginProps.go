package awscdk


// Properties for configuring the CloudFormationValidatePlugin.
//
// Example:
//   // Rules text, read from disk perhaps
//   var myRules string
//   app := awscdk.NewApp()
//
//   awscdk.Validations_Of(app).AddPlugins(awscdk.NewCloudFormationValidatePlugin(&CloudFormationValidatePluginProps{
//   	GuardRules: []ValidationRuleSource{
//   		&ValidationRuleSource{
//   			Name: jsii.String("My rules"),
//   			Content: myRules,
//   		},
//   	},
//   }))
//
type CloudFormationValidatePluginProps struct {
	// Custom Guard rules to evaluate in addition to built-in rules.
	// Default: - no guard rules.
	//
	GuardRules *[]*ValidationRuleSource `field:"optional" json:"guardRules" yaml:"guardRules"`
	// Custom Rego rules to evaluate in addition to built-in rules.
	// Default: - no custom rules.
	//
	RegoRules *[]*ValidationRuleSource `field:"optional" json:"regoRules" yaml:"regoRules"`
}

