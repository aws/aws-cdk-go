package mixinsawsdatazone


// The parameters of the console link specified as part of the environment action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsConsoleLinkParametersProperty := &AwsConsoleLinkParametersProperty{
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentactions-awsconsolelinkparameters.html
//
type CfnEnvironmentActionsPropsMixin_AwsConsoleLinkParametersProperty struct {
	// The URI of the console link specified as part of the environment action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentactions-awsconsolelinkparameters.html#cfn-datazone-environmentactions-awsconsolelinkparameters-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

