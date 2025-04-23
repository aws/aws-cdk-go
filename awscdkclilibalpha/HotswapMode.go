package awscdkclilibalpha


// Experimental.
type HotswapMode string

const (
	// Will fall back to CloudFormation when a non-hotswappable change is detected.
	// Experimental.
	HotswapMode_FALL_BACK HotswapMode = "FALL_BACK"
	// Will not fall back to CloudFormation when a non-hotswappable change is detected.
	// Experimental.
	HotswapMode_HOTSWAP_ONLY HotswapMode = "HOTSWAP_ONLY"
	// Will not attempt to hotswap anything and instead go straight to CloudFormation.
	// Experimental.
	HotswapMode_FULL_DEPLOYMENT HotswapMode = "FULL_DEPLOYMENT"
)

