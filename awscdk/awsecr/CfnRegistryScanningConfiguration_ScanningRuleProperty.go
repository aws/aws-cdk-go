package awsecr


// The scanning rules associated with the registry.
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
	// The details of a scanning repository filter.
	//
	// For more information on how to use filters, see [Using filters](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html#image-scanning-filters) in the *Amazon Elastic Container Registry User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-scanningrule.html#cfn-ecr-registryscanningconfiguration-scanningrule-repositoryfilters
	//
	RepositoryFilters interface{} `field:"required" json:"repositoryFilters" yaml:"repositoryFilters"`
	// The frequency that scans are performed at for a private registry.
	//
	// When the `ENHANCED` scan type is specified, the supported scan frequencies are `CONTINUOUS_SCAN` and `SCAN_ON_PUSH` . When the `BASIC` scan type is specified, the `SCAN_ON_PUSH` scan frequency is supported. If scan on push is not specified, then the `MANUAL` scan frequency is set by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-scanningrule.html#cfn-ecr-registryscanningconfiguration-scanningrule-scanfrequency
	//
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

