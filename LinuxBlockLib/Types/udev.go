package types

type UdevData struct {
	DevID       string            `json:"DevID"`
	Type        string            `json:"Type" udev:"ID_TYPE"`
	Bus         string            `json:"Bus" udev:"BUS"`
	Model       string            `json:"Model" udev:"ID_MODEL"`
	ModelEnc    string            `json:"ModelEncoded" udev:"ID_MODEL_ENC"`
	SerialShort string            `json:"SerialShort" udev:"SERIAL_SHORT"`
	Serial      string            `json:"Serial" udev:"SERIAL"`
	WWN         string            `json:"WWN" udev:"ID_WWN"`
	WWNExtra    string            `json:"WWNWithExtensions" udev:"ID_WWN_WITH_EXTENSION"`
	Path        string            `json:"Path" udev:"ID_PATH"`
	ATAFeature  *ATASmartFeatures `json:"ATASmartFeatures"`
	Raw         string            `json:"RawData"`
}

type ATASmartFeatures struct {
	SecuritySupport bool `json:"SecuritySupport" udev:"ID_ATA_FEATURE_SET_SECURITY"`
	SecurityEnabled bool `json:"SecurityEnabled" udev:"ID_ATA_FEATURE_SET_SECURITY_ENABLED"`
	SmartSupport    bool `json:"SmartSupport" udev:"ID_ATA_FEATURE_SET_SMART"`
	SmartEnabled    bool `json:"SmartEnabled" udev:"ID_ATA_FEATURE_SET_SMART_ENABLED"`
}
