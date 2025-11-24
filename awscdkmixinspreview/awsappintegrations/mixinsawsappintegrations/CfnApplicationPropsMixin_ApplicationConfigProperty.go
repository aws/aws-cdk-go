package mixinsawsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationConfigProperty := &ApplicationConfigProperty{
//   	ContactHandling: &ContactHandlingProperty{
//   		Scope: jsii.String("scope"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationconfig.html
//
type CfnApplicationPropsMixin_ApplicationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationconfig.html#cfn-appintegrations-application-applicationconfig-contacthandling
	//
	ContactHandling interface{} `field:"optional" json:"contactHandling" yaml:"contactHandling"`
}

