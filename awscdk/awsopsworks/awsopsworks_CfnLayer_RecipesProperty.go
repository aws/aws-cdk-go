package awsopsworks


// AWS OpsWorks Stacks supports five lifecycle events: *setup* , *configuration* , *deploy* , *undeploy* , and *shutdown* .
//
// For each layer, AWS OpsWorks Stacks runs a set of standard recipes for each event. In addition, you can provide custom recipes for any or all layers and events. AWS OpsWorks Stacks runs custom event recipes after the standard recipes. `LayerCustomRecipes` specifies the custom recipes for a particular layer to be run in response to each of the five events.
//
// To specify a recipe, use the cookbook's directory name in the repository followed by two colons and the recipe name, which is the recipe's file name without the .rb extension. For example: phpapp2::dbsetup specifies the dbsetup.rb recipe in the repository's phpapp2 folder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipesProperty := &recipesProperty{
//   	configure: []*string{
//   		jsii.String("configure"),
//   	},
//   	deploy: []*string{
//   		jsii.String("deploy"),
//   	},
//   	setup: []*string{
//   		jsii.String("setup"),
//   	},
//   	shutdown: []*string{
//   		jsii.String("shutdown"),
//   	},
//   	undeploy: []*string{
//   		jsii.String("undeploy"),
//   	},
//   }
//
type CfnLayer_RecipesProperty struct {
	// An array of custom recipe names to be run following a `configure` event.
	Configure *[]*string `field:"optional" json:"configure" yaml:"configure"`
	// An array of custom recipe names to be run following a `deploy` event.
	Deploy *[]*string `field:"optional" json:"deploy" yaml:"deploy"`
	// An array of custom recipe names to be run following a `setup` event.
	Setup *[]*string `field:"optional" json:"setup" yaml:"setup"`
	// An array of custom recipe names to be run following a `shutdown` event.
	Shutdown *[]*string `field:"optional" json:"shutdown" yaml:"shutdown"`
	// An array of custom recipe names to be run following a `undeploy` event.
	Undeploy *[]*string `field:"optional" json:"undeploy" yaml:"undeploy"`
}

