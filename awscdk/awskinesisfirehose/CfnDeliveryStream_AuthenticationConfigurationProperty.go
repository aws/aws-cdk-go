package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationProperty := &AuthenticationConfigurationProperty{
//   	Connectivity: jsii.String("connectivity"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-authenticationconfiguration.html
//
type CfnDeliveryStream_AuthenticationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-authenticationconfiguration.html#cfn-kinesisfirehose-deliverystream-authenticationconfiguration-connectivity
	//
	Connectivity *string `field:"required" json:"connectivity" yaml:"connectivity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-authenticationconfiguration.html#cfn-kinesisfirehose-deliverystream-authenticationconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

