package previewawscustomerprofilesmixins


// Configuration and status of the data store for the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreProperty := &DataStoreProperty{
//   	Enabled: jsii.Boolean(false),
//   	Readiness: &ReadinessProperty{
//   		Message: jsii.String("message"),
//   		ProgressPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-datastore.html
//
type CfnDomainPropsMixin_DataStoreProperty struct {
	// Whether the data store is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-datastore.html#cfn-customerprofiles-domain-datastore-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Progress information for data store setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-datastore.html#cfn-customerprofiles-domain-datastore-readiness
	//
	Readiness interface{} `field:"optional" json:"readiness" yaml:"readiness"`
}

