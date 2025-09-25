package awspcaconnectorad


// General flags for v3 template schema that defines if the template is for a machine or a user and if the template can be issued using autoenrollment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalFlagsV3Property := &GeneralFlagsV3Property{
//   	AutoEnrollment: jsii.Boolean(false),
//   	MachineType: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv3.html
//
type CfnTemplate_GeneralFlagsV3Property struct {
	// Allows certificate issuance using autoenrollment.
	//
	// Set to TRUE to allow autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv3.html#cfn-pcaconnectorad-template-generalflagsv3-autoenrollment
	//
	AutoEnrollment interface{} `field:"optional" json:"autoEnrollment" yaml:"autoEnrollment"`
	// Defines if the template is for machines or users.
	//
	// Set to TRUE if the template is for machines. Set to FALSE if the template is for users
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv3.html#cfn-pcaconnectorad-template-generalflagsv3-machinetype
	//
	MachineType interface{} `field:"optional" json:"machineType" yaml:"machineType"`
}

