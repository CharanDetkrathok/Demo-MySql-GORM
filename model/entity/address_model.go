package entity

type (
	District struct {
		DISTRICTID   int    `gorm:"column:DISTRICT_ID;NOT NULL"`
		DISTRICTCODE string `gorm:"column:DISTRICT_CODE;NOT NULL"`
		DISTRICTNAME string `gorm:"column:DISTRICT_NAME;NOT NULL"`
		GEOID        int    `gorm:"column:GEO_ID;NOT NULL"`
		PROVINCEID   int    `gorm:"column:PROVINCE_ID;NOT NULL"`
	}

	Geography struct {
		GEOID   int    `gorm:"column:GEO_ID;AUTO_INCREMENT;NOT NULL"`
		GEONAME string `gorm:"column:GEO_NAME;NOT NULL"`
	}

	Province struct {
		PROVINCEID   int    `gorm:"column:PROVINCE_ID;NOT NULL"`
		PROVINCECODE string `gorm:"column:PROVINCE_CODE;NOT NULL"`
		PROVINCENAME string `gorm:"column:PROVINCE_NAME;NOT NULL"`
		GEOID        int    `gorm:"column:GEO_ID;NOT NULL"`
	}

	SubDistrict struct {
		SUBDISTRICTID   int    `gorm:"column:SUB_DISTRICT_ID;NOT NULL"`
		SUBDISTRICTCODE string `gorm:"column:SUB_DISTRICT_CODE;NOT NULL"`
		SUBDISTRICTNAME string `gorm:"column:SUB_DISTRICT_NAME;NOT NULL"`
		DISTRICTID      int    `gorm:"column:DISTRICT_ID;NOT NULL"`
		PROVINCEID      int    `gorm:"column:PROVINCE_ID;NOT NULL"`
		GEOID           int    `gorm:"column:GEO_ID;NOT NULL"`
	}

	Zipcode struct {
		ZIPCODEID       int    `gorm:"column:ZIPCODE_ID;NOT NULL"`
		SUBDISTRICTCODE string `gorm:"column:SUB_DISTRICT_CODE;NOT NULL"`
		PROVINCEID      string `gorm:"column:PROVINCE_ID;NOT NULL"`
		DISTRICTID      string `gorm:"column:DISTRICT_ID;NOT NULL"`
		SUBDISTRICTID   string `gorm:"column:SUB_DISTRICT_ID;NOT NULL"`
		ZIPCODE         string `gorm:"column:ZIPCODE;NOT NULL"`
	}
)

func (*District) TableName() string {
	return "district"
}

func (*Geography) TableName() string {
	return "geography"
}

func (*Province) TableName() string {
	return "province"
}

func (*SubDistrict) TableName() string {
	return "subDistrict"
}

func (*Zipcode) TableName() string {
	return "zipcode"
}
