package previewawsdevopsagentmixins


// Dynatrace service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registeredDynatraceDetailsProperty := &RegisteredDynatraceDetailsProperty{
//   	AccountUrn: jsii.String("accountUrn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registereddynatracedetails.html
//
type CfnServicePropsMixin_RegisteredDynatraceDetailsProperty struct {
	// Dynatrace resource account URN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registereddynatracedetails.html#cfn-devopsagent-service-registereddynatracedetails-accounturn
	//
	AccountUrn *string `field:"optional" json:"accountUrn" yaml:"accountUrn"`
}

