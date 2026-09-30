package awsmgn


// An AWS Systems Manager Document external parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ssmExternalParameterProperty := &SsmExternalParameterProperty{
//   	DynamicPath: jsii.String("dynamicPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmexternalparameter.html
//
type CfnLaunchConfigurationTemplatePropsMixin_SsmExternalParameterProperty struct {
	// AWS Systems Manager Document external parameter dynamic path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmexternalparameter.html#cfn-mgn-launchconfigurationtemplate-ssmexternalparameter-dynamicpath
	//
	DynamicPath *string `field:"optional" json:"dynamicPath" yaml:"dynamicPath"`
}

