package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPGSV prefix of GPGSV sentence type
	PrefixGPGSV = "GPGSV"
)

// GPGSV represents overview satellite view (SV).
// http://aprs.gids.nl/nmea/#gsa
type GPGSV struct {
	Sentence
	// Total number of messages of this type in this cycle.
	TotalNumberOfMessages int64
	// Message number
	NumberOFMessage int64
	// Total number of SVs in view
	TotalNumberOfSVs int64
	// List of satellite view information
	//8-11 = Information about second SV, same as field 4-7
	//12-15= Information about third SV, same as field 4-7
	//16-19= Information about fourth SV, same as field 4-7
	SVList []SatelliteView
}

type SatelliteView struct {
	// SV PRN number
	SVPRNNumber string
	// Elevation in degrees, 90 maximum
	ElevationInDegrees string
	// Azimuth, degrees from true north, 000 to 359
	Azimuth string
	//SNR, 00-99 dB (null when not tracking)
	SNR string
}

// NewGPGSV constructor
func NewGPGSV(sentence Sentence) GPGSV {
	s := new(GPGSV)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPGSV) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPGSV sentence into this struct.
func (s *GPGSV) parse() error {
	if s.Type != PrefixGPGSV {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPGSV)
	}

	s.TotalNumberOfMessages, _ = strconv.ParseInt(s.Fields[0], 10, 64)
	s.NumberOFMessage, _ = strconv.ParseInt(s.Fields[1], 10, 64)
	s.TotalNumberOfSVs, _ = strconv.ParseInt(s.Fields[2], 10, 64)

	// Satellites in view.
	s.SVList = make([]SatelliteView, 0, 4)
	for i := 3; i < 19; i += 4 {
		sv := new(SatelliteView)
		sv.SVPRNNumber = s.Fields[i]
		sv.ElevationInDegrees = s.Fields[i+1]
		sv.Azimuth = s.Fields[i+2]
		sv.SNR = s.Fields[i+3]
		s.SVList = append(s.SVList, *sv)
	}

	return nil
}
