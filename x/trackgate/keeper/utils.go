package keeper

import (
	"fmt"
	"regexp"
	"time"
)

// ValidateTrackID checks if the given trackId follows the pattern TRK-yyyyMMdd-nnnn
func ValidateTrackID(trackId string) bool {
	// Create a regular expression to match the track ID format
	date := time.Now().Format("20060102")
	regexPattern := fmt.Sprintf("^TRK-%s-[0-9]{4}$", date)
	match, _ := regexp.MatchString(regexPattern, trackId)
	return match
}
