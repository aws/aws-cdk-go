package previewawspcaconnectoradmixins


// General flags for v2 template schema that defines if the template is for a machine or a user and if the template can be issued using autoenrollment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   generalFlagsV2Property := &GeneralFlagsV2Property{
//   	AutoEnrollment: jsii.Boolean(false),
//   	MachineType: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv2.html
//
type CfnTemplatePropsMixin_GeneralFlagsV2Property struct {
	// Allows certificate issuance using autoenrollment.
	//
	// Set to TRUE to allow autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv2.html#cfn-pcaconnectorad-template-generalflagsv2-autoenrollment
	//
	AutoEnrollment interface{} `field:"optional" json:"autoEnrollment" yaml:"autoEnrollment"`
	// Defines if the template is for machines or users.
	//
	// Set to TRUE if the template is for machines. Set to FALSE if the template is for users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv2.html#cfn-pcaconnectorad-template-generalflagsv2-machinetype
	//
	MachineType interface{} `field:"optional" json:"machineType" yaml:"machineType"`
}

