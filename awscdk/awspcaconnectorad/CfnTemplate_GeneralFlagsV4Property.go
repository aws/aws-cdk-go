package awspcaconnectorad


// General flags for v4 template schema that defines if the template is for a machine or a user and if the template can be issued using autoenrollment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalFlagsV4Property := &GeneralFlagsV4Property{
//   	AutoEnrollment: jsii.Boolean(false),
//   	MachineType: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv4.html
//
type CfnTemplate_GeneralFlagsV4Property struct {
	// Allows certificate issuance using autoenrollment.
	//
	// Set to TRUE to allow autoenrollment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv4.html#cfn-pcaconnectorad-template-generalflagsv4-autoenrollment
	//
	AutoEnrollment interface{} `field:"optional" json:"autoEnrollment" yaml:"autoEnrollment"`
	// Defines if the template is for machines or users.
	//
	// Set to TRUE if the template is for machines. Set to FALSE if the template is for users
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-generalflagsv4.html#cfn-pcaconnectorad-template-generalflagsv4-machinetype
	//
	MachineType interface{} `field:"optional" json:"machineType" yaml:"machineType"`
}

