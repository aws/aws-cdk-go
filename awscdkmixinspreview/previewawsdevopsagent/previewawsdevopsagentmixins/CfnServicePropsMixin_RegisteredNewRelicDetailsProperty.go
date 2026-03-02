package previewawsdevopsagentmixins


// New Relic service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registeredNewRelicDetailsProperty := &RegisteredNewRelicDetailsProperty{
//   	AccountId: jsii.String("accountId"),
//   	Description: jsii.String("description"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html
//
type CfnServicePropsMixin_RegisteredNewRelicDetailsProperty struct {
	// New Relic account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Optional user description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// New Relic region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

