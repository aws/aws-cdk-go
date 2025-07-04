package awsiot


// Properties for defining a `CfnThingPrincipalAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThingPrincipalAttachmentProps := &CfnThingPrincipalAttachmentProps{
//   	Principal: jsii.String("principal"),
//   	ThingName: jsii.String("thingName"),
//
//   	// the properties below are optional
//   	ThingPrincipalType: jsii.String("thingPrincipalType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
//
type CfnThingPrincipalAttachmentProps struct {
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The name of the AWS IoT thing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingname
	//
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingprincipaltype
	//
	ThingPrincipalType *string `field:"optional" json:"thingPrincipalType" yaml:"thingPrincipalType"`
}

