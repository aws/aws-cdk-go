package previewawscodebuildevents


// Type definition for Vpc-config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConfig := &VpcConfig{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Subnets: []VpcConfigItem{
//   		&VpcConfigItem{
//   			BuildFleetAz: []*string{
//   				jsii.String("buildFleetAz"),
//   			},
//   			CustomerAz: []*string{
//   				jsii.String("customerAz"),
//   			},
//   			SubnetId: []*string{
//   				jsii.String("subnetId"),
//   			},
//   		},
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_VpcConfig struct {
	// security-group-ids property.
	//
	// Specify an array of string values to match this event if the actual value of security-group-ids is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// subnets property.
	//
	// Specify an array of string values to match this event if the actual value of subnets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Subnets *[]*ProjectEvents_CodeBuildBuildPhaseChange_VpcConfigItem `field:"optional" json:"subnets" yaml:"subnets"`
	// vpc-id property.
	//
	// Specify an array of string values to match this event if the actual value of vpc-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

