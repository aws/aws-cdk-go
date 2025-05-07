package awsecr


// Properties for defining a `CfnRegistryScanningConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegistryScanningConfigurationProps := &CfnRegistryScanningConfigurationProps{
//   	Rules: []interface{}{
//   		&ScanningRuleProperty{
//   			RepositoryFilters: []interface{}{
//   				&RepositoryFilterProperty{
//   					Filter: jsii.String("filter"),
//   					FilterType: jsii.String("filterType"),
//   				},
//   			},
//   			ScanFrequency: jsii.String("scanFrequency"),
//   		},
//   	},
//   	ScanType: jsii.String("scanType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html
//
type CfnRegistryScanningConfigurationProps struct {
	// The scanning rules associated with the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html#cfn-ecr-registryscanningconfiguration-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The type of scanning configured for the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registryscanningconfiguration.html#cfn-ecr-registryscanningconfiguration-scantype
	//
	ScanType *string `field:"required" json:"scanType" yaml:"scanType"`
}

