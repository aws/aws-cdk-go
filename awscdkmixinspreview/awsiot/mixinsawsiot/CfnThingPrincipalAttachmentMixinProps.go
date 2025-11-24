package mixinsawsiot


// Properties for CfnThingPrincipalAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThingPrincipalAttachmentMixinProps := &CfnThingPrincipalAttachmentMixinProps{
//   	Principal: jsii.String("principal"),
//   	ThingName: jsii.String("thingName"),
//   	ThingPrincipalType: jsii.String("thingPrincipalType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
//
type CfnThingPrincipalAttachmentMixinProps struct {
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// The name of the AWS IoT thing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingname
	//
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingprincipaltype
	//
	ThingPrincipalType *string `field:"optional" json:"thingPrincipalType" yaml:"thingPrincipalType"`
}

