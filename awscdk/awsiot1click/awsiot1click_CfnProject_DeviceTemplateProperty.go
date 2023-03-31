package awsiot1click


// In AWS CloudFormation , use the `DeviceTemplate` property type to define the template for an AWS IoT 1-Click project.
//
// `DeviceTemplate` is a property of the `AWS::IoT1Click::Project` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var callbackOverrides interface{}
//
//   deviceTemplateProperty := &deviceTemplateProperty{
//   	callbackOverrides: callbackOverrides,
//   	deviceType: jsii.String("deviceType"),
//   }
//
type CfnProject_DeviceTemplateProperty struct {
	// An optional AWS Lambda function to invoke instead of the default AWS Lambda function provided by the placement template.
	CallbackOverrides interface{} `field:"optional" json:"callbackOverrides" yaml:"callbackOverrides"`
	// The device type, which currently must be `"button"` .
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

