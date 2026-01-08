package previewawscodebuildevents


// Type definition for Vpc-configItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConfigItem := &VpcConfigItem{
//   	BuildFleetAz: []*string{
//   		jsii.String("buildFleetAz"),
//   	},
//   	CustomerAz: []*string{
//   		jsii.String("customerAz"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_VpcConfigItem struct {
	// build-fleet-az property.
	//
	// Specify an array of string values to match this event if the actual value of build-fleet-az is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BuildFleetAz *[]*string `field:"optional" json:"buildFleetAz" yaml:"buildFleetAz"`
	// customer-az property.
	//
	// Specify an array of string values to match this event if the actual value of customer-az is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CustomerAz *[]*string `field:"optional" json:"customerAz" yaml:"customerAz"`
	// subnet-id property.
	//
	// Specify an array of string values to match this event if the actual value of subnet-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

