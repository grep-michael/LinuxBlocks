package types

type UdevData struct {
	DevID       string           `json:"DevID"`
	Type        string           `json:"Type" udev:"E:ID_TYPE"`
	Bus         string           `json:"Bus" udev:"E:BUS"`
	Model       string           `json:"Model" udev:"E:ID_MODEL"`
	ModelEnc    string           `json:"ModelEncoded" udev:"E:ID_MODEL_ENC"`
	SerialShort string           `json:"SerialShort" udev:"E:SERIAL_SHORT"`
	Serial      string           `json:"Serial" udev:"E:SERIAL"`
	WWN         string           `json:"WWN" udev:"E:ID_WWN"`
	WWNExtra    string           `json:"WWNWithExtensions" udev:"E:ID_WWN_WITH_EXTENSION"`
	Path        string           `json:"Path" udev:"E:ID_PATH"`
	ATAFeature  ATASmartFeatures `json:"ATASmartFeatures"`
	Raw         string           `json:"RawData"`
}

type ATASmartFeatures struct {
	SecuritySupport bool `json:"SecuritySupport" udev:"E:ID_ATA_FEATURE_SET_SECURITY"`
	SecurityEnabled bool `json:"SecurityEnabled" udev:"D_ATA_FEATURE_SET_SECURITY_ENABLED"`
	SmartSupport    bool `json:"SmartSupport" udev:"ID_ATA_FEATURE_SET_SMART"`
	SmartEnabled    bool `json:"SmartEnabled" udev:"ID_ATA_FEATURE_SET_SMART_ENABLED"`
}
