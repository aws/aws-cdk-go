package awsses


// Scaling mode to use for this IP pool.
//
// Example:
//   ses.NewDedicatedIpPool(this, jsii.String("Pool"), &DedicatedIpPoolProps{
//   	DedicatedIpPoolName: jsii.String("mypool"),
//   	ScalingMode: ses.ScalingMode_STANDARD,
//   })
//
// See: https://docs.aws.amazon.com/ses/latest/dg/dedicated-ip.html
//
type ScalingMode string

const (
	// The customer controls which IPs are part of the dedicated IP pool.
	ScalingMode_STANDARD ScalingMode = "STANDARD"
	// The reputation and number of IPs are automatically managed by Amazon SES.
	ScalingMode_MANAGED ScalingMode = "MANAGED"
)

