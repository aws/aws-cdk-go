package previewawsdevopsagentmixins


// ServiceNow service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registeredServiceNowDetailsProperty := &RegisteredServiceNowDetailsProperty{
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredservicenowdetails.html
//
type CfnServicePropsMixin_RegisteredServiceNowDetailsProperty struct {
	// ServiceNow instance URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredservicenowdetails.html#cfn-devopsagent-service-registeredservicenowdetails-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

