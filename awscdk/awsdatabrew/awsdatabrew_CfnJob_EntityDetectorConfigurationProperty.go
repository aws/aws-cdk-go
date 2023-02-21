package awsdatabrew


// Configuration of entity detection for a profile job.
//
// When undefined, entity detection is disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityDetectorConfigurationProperty := &entityDetectorConfigurationProperty{
//   	entityTypes: []*string{
//   		jsii.String("entityTypes"),
//   	},
//
//   	// the properties below are optional
//   	allowedStatistics: &allowedStatisticsProperty{
//   		statistics: []*string{
//   			jsii.String("statistics"),
//   		},
//   	},
//   }
//
type CfnJob_EntityDetectorConfigurationProperty struct {
	// Entity types to detect. Can be any of the following:.
	//
	// - USA_SSN
	// - EMAIL
	// - USA_ITIN
	// - USA_PASSPORT_NUMBER
	// - PHONE_NUMBER
	// - USA_DRIVING_LICENSE
	// - BANK_ACCOUNT
	// - CREDIT_CARD
	// - IP_ADDRESS
	// - MAC_ADDRESS
	// - USA_DEA_NUMBER
	// - USA_HCPCS_CODE
	// - USA_NATIONAL_PROVIDER_IDENTIFIER
	// - USA_NATIONAL_DRUG_CODE
	// - USA_HEALTH_INSURANCE_CLAIM_NUMBER
	// - USA_MEDICARE_BENEFICIARY_IDENTIFIER
	// - USA_CPT_CODE
	// - PERSON_NAME
	// - DATE
	//
	// The Entity type group USA_ALL is also supported, and includes all of the above entity types except PERSON_NAME and DATE.
	EntityTypes *[]*string `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Configuration of statistics that are allowed to be run on columns that contain detected entities.
	//
	// When undefined, no statistics will be computed on columns that contain detected entities.
	AllowedStatistics interface{} `field:"optional" json:"allowedStatistics" yaml:"allowedStatistics"`
}

