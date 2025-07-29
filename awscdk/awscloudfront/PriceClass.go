package awscloudfront


// The price class determines how many edge locations CloudFront will use for your distribution.
//
// See https://aws.amazon.com/cloudfront/pricing/ for full list of supported regions.
type PriceClass string

const (
	// USA, Canada, Europe, & Israel.
	PriceClass_PRICE_CLASS_100 PriceClass = "PRICE_CLASS_100"
	// PRICE_CLASS_100 + South Africa, Kenya, Middle East, Japan, Singapore, South Korea, Taiwan, Hong Kong, & Philippines.
	PriceClass_PRICE_CLASS_200 PriceClass = "PRICE_CLASS_200"
	// All locations.
	PriceClass_PRICE_CLASS_ALL PriceClass = "PRICE_CLASS_ALL"
)

