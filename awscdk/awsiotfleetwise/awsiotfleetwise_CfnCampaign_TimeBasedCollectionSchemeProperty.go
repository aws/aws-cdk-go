package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedCollectionSchemeProperty := &timeBasedCollectionSchemeProperty{
//   	periodMs: jsii.Number(123),
//   }
//
type CfnCampaign_TimeBasedCollectionSchemeProperty struct {
	// `CfnCampaign.TimeBasedCollectionSchemeProperty.PeriodMs`.
	PeriodMs *float64 `field:"required" json:"periodMs" yaml:"periodMs"`
}

