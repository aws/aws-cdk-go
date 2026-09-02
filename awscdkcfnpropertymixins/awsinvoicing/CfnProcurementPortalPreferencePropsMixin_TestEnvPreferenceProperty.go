package awsinvoicing


// Configuration settings for the test environment of the procurement portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   testEnvPreferenceProperty := &TestEnvPreferenceProperty{
//   	BuyerDomain: jsii.String("buyerDomain"),
//   	BuyerIdentifier: jsii.String("buyerIdentifier"),
//   	ProcurementPortalInstanceEndpoint: jsii.String("procurementPortalInstanceEndpoint"),
//   	ProcurementPortalSharedSecret: jsii.String("procurementPortalSharedSecret"),
//   	SupplierDomain: jsii.String("supplierDomain"),
//   	SupplierIdentifier: jsii.String("supplierIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html
//
type CfnProcurementPortalPreferencePropsMixin_TestEnvPreferenceProperty struct {
	// The domain identifier for the buyer in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-buyerdomain
	//
	BuyerDomain *string `field:"optional" json:"buyerDomain" yaml:"buyerDomain"`
	// The unique identifier for the buyer in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-buyeridentifier
	//
	BuyerIdentifier *string `field:"optional" json:"buyerIdentifier" yaml:"buyerIdentifier"`
	// The endpoint URL for e-invoice delivery in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-procurementportalinstanceendpoint
	//
	ProcurementPortalInstanceEndpoint *string `field:"optional" json:"procurementPortalInstanceEndpoint" yaml:"procurementPortalInstanceEndpoint"`
	// The shared secret for secure communication in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-procurementportalsharedsecret
	//
	ProcurementPortalSharedSecret *string `field:"optional" json:"procurementPortalSharedSecret" yaml:"procurementPortalSharedSecret"`
	// The domain identifier for the supplier in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-supplierdomain
	//
	SupplierDomain *string `field:"optional" json:"supplierDomain" yaml:"supplierDomain"`
	// The unique identifier for the supplier in the test environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-testenvpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference-supplieridentifier
	//
	SupplierIdentifier *string `field:"optional" json:"supplierIdentifier" yaml:"supplierIdentifier"`
}

