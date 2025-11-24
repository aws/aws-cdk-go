package mixinsawsopsworks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recipesProperty := &RecipesProperty{
//   	Configure: []*string{
//   		jsii.String("configure"),
//   	},
//   	Deploy: []*string{
//   		jsii.String("deploy"),
//   	},
//   	Setup: []*string{
//   		jsii.String("setup"),
//   	},
//   	Shutdown: []*string{
//   		jsii.String("shutdown"),
//   	},
//   	Undeploy: []*string{
//   		jsii.String("undeploy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
//
type CfnLayerPropsMixin_RecipesProperty struct {
	// An array of custom recipe names to be run following a `configure` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-recipes-configure
	//
	Configure *[]*string `field:"optional" json:"configure" yaml:"configure"`
	// An array of custom recipe names to be run following a `deploy` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-recipes-deploy
	//
	Deploy *[]*string `field:"optional" json:"deploy" yaml:"deploy"`
	// An array of custom recipe names to be run following a `setup` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-recipes-setup
	//
	Setup *[]*string `field:"optional" json:"setup" yaml:"setup"`
	// An array of custom recipe names to be run following a `shutdown` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-recipes-shutdown
	//
	Shutdown *[]*string `field:"optional" json:"shutdown" yaml:"shutdown"`
	// An array of custom recipe names to be run following a `undeploy` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-recipes-undeploy
	//
	Undeploy *[]*string `field:"optional" json:"undeploy" yaml:"undeploy"`
}

