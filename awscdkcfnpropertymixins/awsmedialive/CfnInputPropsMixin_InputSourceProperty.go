package awsmedialive


// The settings for a PULL type input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inputSourceProperty := &InputSourceProperty{
//   	PasswordParam: jsii.String("passwordParam"),
//   	Url: jsii.String("url"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsource.html
//
type CfnInputPropsMixin_InputSourceProperty struct {
	// The key used to extract the password from EC2 Parameter store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsource.html#cfn-medialive-input-inputsource-passwordparam
	//
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// This represents the customer's source URL where stream is pulled from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsource.html#cfn-medialive-input-inputsource-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The username for the input source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputsource.html#cfn-medialive-input-inputsource-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

