package awsecr


// A rule representing the details of a scanning configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scanningRuleProperty := &ScanningRuleProperty{
//   	RepositoryFilters: []interface{}{
//   		&RepositoryFilterProperty{
//   			Filter: jsii.String("filter"),
//   			FilterType: jsii.String("filterType"),
//   		},
//   	},
//   	ScanFrequency: jsii.String("scanFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-scanningrule.html
//
type CfnRegistryScanningConfiguration_ScanningRuleProperty struct {
	// The repository filters associated with the scanning configuration for a private registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-scanningrule.html#cfn-ecr-registryscanningconfiguration-scanningrule-repositoryfilters
	//
	RepositoryFilters interface{} `field:"required" json:"repositoryFilters" yaml:"repositoryFilters"`
	// The frequency that scans are performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-scanningrule.html#cfn-ecr-registryscanningconfiguration-scanningrule-scanfrequency
	//
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

