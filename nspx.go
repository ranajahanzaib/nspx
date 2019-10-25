package nspx

// TYPES Declaration

// Product type
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	ProductType   string `json:"type"`
	ProductPrice  string `json:"price"`
	ProductWeight int    `json:"productWeight"`
	BatteryAvl    bool   `json:"batteryAvl"`
	BatteryType   string `json:"batteryType"`
	BatteryPower  int    `json:"batteryPower"`
//	color, brand,
}

// Electronics type
// self powered / powered electronics
// screenSize, batteryPower, batteryTime, fingerPrint (self-powered)
// cameras, wifi, blueTooth, chargerType, audioJack,
// monitors, storages (hardDisk, usb, ssd), headPhones, otherSoundProducts, ram, cpu, pci, graphicCard,
// 
type Electronics struct {
}

// Laptop type
type Laptop struct {
	BacklitKb         bool   `json:"backlitKb"`
	TouchScreen       bool   `json:"touchScreen"`
	WiFi              bool   `json:"wifi"`
	WiFiConnectivity  string `json:"wifiConnectivity"`
	ScreenToBodyRatio int    `json:"screenToBodyRatio"`
	CPU               string `json:"cpu"`
	GPU               string `json:"gpu"`
	DisplaySize       int    `json:"displaySize"`
	DisplayResolution string `json:"displayResolution"`
	RAMSticks         int    `json:"ramSticks"`
	RAMSize           int    `json:"ramSize"`
	DDRGen            int    `json:"ddrGen"`
}

// Pc type
type Pc struct {
}

// Mac type
type Mac struct {
}

// Mobile type
type Mobile struct {
}

// LaptopParts type
type LaptopParts struct {
}

// PcParts type
type PcParts struct {
}

// MobileParts type
type MobileParts struct {
}
