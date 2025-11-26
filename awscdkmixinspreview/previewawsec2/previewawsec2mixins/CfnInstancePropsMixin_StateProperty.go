package previewawsec2mixins


// Describes the current state of an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stateProperty := &StateProperty{
//   	Code: jsii.String("code"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html
//
type CfnInstancePropsMixin_StateProperty struct {
	// The state of the instance as a 16-bit unsigned integer.
	//
	// The high byte is all of the bits between 2^8 and (2^16)-1, which equals decimal values between 256 and 65,535. These numerical values are used for internal purposes and should be ignored.
	//
	// The low byte is all of the bits between 2^0 and (2^8)-1, which equals decimal values between 0 and 255.
	//
	// The valid values for instance-state-code will all be in the range of the low byte and they are:
	//
	// - `0` : `pending`
	// - `16` : `running`
	// - `32` : `shutting-down`
	// - `48` : `terminated`
	// - `64` : `stopping`
	// - `80` : `stopped`
	//
	// You can ignore the high byte value by zeroing out all of the bits above 2^8 or 256 in decimal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html#cfn-ec2-instance-state-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The current state of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-state.html#cfn-ec2-instance-state-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

