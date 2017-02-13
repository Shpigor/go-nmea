package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//$GPGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,00,13,06,292,00*74
//$GPGSV,3,2,11,14,25,170,00,16,57,208,39,18,67,296,40,19,40,246,00*74
//$GPGSV,3,3,11,22,42,067,42,24,14,311,43,27,05,244,00,,,,*4D
//$GPGSV,1,1,13,02,02,213,,03,-3,000,,11,00,121,,14,13,172,05*67
func TestGPGSVGoodSentence(t *testing.T) {
	goodMsg := "$GPGSV,3,3,11,22,42,067,42,24,14,311,43,27,05,244,00,,,,*4D"
	sentence, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")

	// Attributes of the parsed sentence, and their expected values.
	expected := GPGSV{
		Sentence: Sentence{
			Type:     "GPGSV",
			Fields:   []string{"3", "3", "11", "22", "42", "067", "42", "24", "14", "311", "43", "27", "05", "244", "00", "", "", "", ""},
			Checksum: "4D",
			Raw:      "$GPGSV,3,3,11,22,42,067,42,24,14,311,43,27,05,244,00,,,,*4D",

		},
		TotalNumberOfMessages: 3,
		NumberOFMessage:       3,
		TotalNumberOfSVs:      11,
		SVList: []SatelliteView{
			{
				SVPRNNumber:        "22",
				ElevationInDegrees: "42",
				Azimuth:            "067",
				SNR:                "42",
			},
			{
				SVPRNNumber:        "24",
				ElevationInDegrees: "14",
				Azimuth:            "311",
				SNR:                "43",
			},
			{
				SVPRNNumber:        "27",
				ElevationInDegrees: "05",
				Azimuth:            "244",
				SNR:                "00",
			},
			{
				SVPRNNumber:        "",
				ElevationInDegrees: "",
				Azimuth:            "",
				SNR:                "",
			},
		},
	}
	assert.EqualValues(t, expected, sentence, "Sentence values do not match")
}
