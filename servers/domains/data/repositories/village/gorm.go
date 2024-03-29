package village

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"bagasunix/kuyngetrip/pkg/errors"
	"bagasunix/kuyngetrip/servers/domains/data/models"
)

type gormProvider struct {
	logger zap.Logger
	db     *gorm.DB
}

func (g *gormProvider) GetByIdWithChannel(ctx context.Context, id int64, result chan models.SingleResult[*models.Village]) {
	result <- g.GetById(ctx, id)
}

func (g *gormProvider) GetById(ctx context.Context, id int64) (result models.SingleResult[*models.Village]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Joins("SubDistrict").First(&result.Value, "id = ?", id).Error)
	return result
}

// GetByKeywordsByCoordinateByRadius implements Repository
func (g *gormProvider) GetByKeywordsByCoordinateByRadius(ctx context.Context, limit int64, keywords string, latitude float64, longitude float64, radius float32) (result models.SliceResult[models.Village], err error) {
	type villageModel struct {
		Village_id             int64
		Village_name           string
		Sub_district_id        int32
		Sub_district_name      string
		Sub_district_latitude  float64
		Sub_district_longitude float64
		City_id                int16
		City_name              string
		City_latitude          float64
		City_longitude         float64
		Province_id            int16
		Province_name          string
		Province_latitude      float64
		Province_longitude     float64
		Country_id             int16
		Country_name           string
		Country_latitude       float64
		Country_longitude      float64
		Village_postal_code    string
		Distance               float64
	}

	var rows []villageModel
	keyWord := fmt.Sprint("%", keywords, "%")
	err = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Raw(QUERY_GET_COORDINATE, sql.Named("keyWord", keyWord), sql.Named("limited", limit), sql.Named("latitude", latitude), sql.Named("longitude", longitude), sql.Named("radius", radius)).Scan(&rows).Error)

	for _, v := range rows {
		coBuild := models.NewCountryBuilder()
		coBuild.SetId(v.Country_id)
		coBuild.SetName(v.Country_name)
		coBuild.SetLatitude(v.Country_latitude)
		coBuild.SetLongitude(v.Country_longitude)

		pBuild := models.NewProvinceBuilder()
		pBuild.SetId(v.Province_id)
		pBuild.SetName(v.Province_name)
		pBuild.SetLatitude(v.Province_latitude)
		pBuild.SetLongitude(v.Province_longitude)
		pBuild.SetCountry(coBuild.Build())

		cBuild := models.NewCityBuilder()
		cBuild.SetId(v.City_id)
		cBuild.SetName(v.City_name)
		cBuild.SetLatitude(v.City_latitude)
		cBuild.SetLongitude(v.City_longitude)
		cBuild.SetProvince(pBuild.Build())

		sBuild := models.NewSubDistrictBuilder()
		sBuild.SetId(v.Sub_district_id)
		sBuild.SetName(v.Sub_district_name)
		sBuild.SetLatitude(v.Sub_district_latitude)
		sBuild.SetLongitude(v.Sub_district_longitude)
		sBuild.SetCity(cBuild.Build())

		vBuild := models.NewVillageBuilder()
		vBuild.SetId(v.Village_id)
		vBuild.SetName(v.Village_name)
		vBuild.SetPostalCodes(v.Village_postal_code)
		vBuild.SetDistance(v.Distance)
		vBuild.SetSubDistrict(sBuild.Build())

		result.Value = append(result.Value, *vBuild.Build())
	}
	return result, err
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "village"
}

// GetByKeywords GetVillageByKeywords implements Repository
func (g *gormProvider) GetByKeywords(ctx context.Context, limit int64, keywords string) (result models.SliceResult[models.Village], err error) {

	type villageModel struct {
		Village_id             int64
		Village_name           string
		Sub_district_id        int32
		Sub_district_name      string
		Sub_district_latitude  float64
		Sub_district_longitude float64
		City_id                int16
		City_name              string
		City_latitude          float64
		City_longitude         float64
		Province_id            int16
		Province_name          string
		Province_latitude      float64
		Province_longitude     float64
		Country_id             int16
		Country_name           string
		Country_latitude       float64
		Country_longitude      float64
		Village_postal_code    string
	}

	var rows []villageModel
	keyWord := fmt.Sprint("%", keywords, "%")
	err = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Raw(QUERY_GET_KEYWORD, sql.Named("keyWord", keyWord), sql.Named("limited", limit)).Scan(&rows).Error)

	for _, v := range rows {
		coBuild := models.NewCountryBuilder()
		coBuild.SetId(v.Country_id)
		coBuild.SetName(v.Country_name)
		coBuild.SetLatitude(v.Country_latitude)
		coBuild.SetLongitude(v.Country_longitude)

		pBuild := models.NewProvinceBuilder()
		pBuild.SetId(v.Province_id)
		pBuild.SetName(v.Province_name)
		pBuild.SetLatitude(v.Province_latitude)
		pBuild.SetLongitude(v.Province_longitude)
		pBuild.SetCountry(coBuild.Build())

		cBuild := models.NewCityBuilder()
		cBuild.SetId(v.City_id)
		cBuild.SetName(v.City_name)
		cBuild.SetLatitude(v.City_latitude)
		cBuild.SetLongitude(v.City_longitude)
		cBuild.SetProvince(pBuild.Build())

		sBuild := models.NewSubDistrictBuilder()
		sBuild.SetId(v.Sub_district_id)
		sBuild.SetName(v.Sub_district_name)
		sBuild.SetLatitude(v.Sub_district_latitude)
		sBuild.SetLongitude(v.Sub_district_longitude)
		sBuild.SetCity(cBuild.Build())

		vBuild := models.NewVillageBuilder()
		vBuild.SetId(v.Village_id)
		vBuild.SetName(v.Village_name)
		vBuild.SetPostalCodes(v.Village_postal_code)
		vBuild.SetSubDistrict(sBuild.Build())

		result.Value = append(result.Value, *vBuild.Build())
	}
	return result, err
}

func NewGorm(logger zap.Logger, db *gorm.DB) Repository {
	a := new(gormProvider)
	a.logger = logger
	a.db = db
	return a
}
