package polyfetcher

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeometryUtils_FetchPolygons_success(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"frankfurt", "munich"}

	//When
	response, err := gm.FetchPolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	fmt.Print(response)
	assert.NotNil(t, response)
}

func TestGeometryUtils_FetchPolygons_Interested_in_only_Polygons_or_Multi_Polygons(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"Rotterdam"}

	//When
	response, err := gm.FetchPolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, response[0].Type, Polygon)
}

func TestGeometryUtils_FetchPolygons_Interested_in_only_Polygons_or_Multi_Polygons_Return_Empty_When_Point(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"Mysore"}

	//When
	response, err := gm.FetchPolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, response[0].Type, "")
}

func TestGeometryUtils_FetchPolygons_one_area_is_invalid(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"asdasdasdasd", "asdasdasdasd"}

	//When
	_, err := gm.FetchPolygons(context.Background(), areas)

	//Then
	assert.NotNil(t, err)
}

func TestGeometryUtils_CombinePolygons_Success(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"frankfurt", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGeometryUtils_CombinePolygons_areas_one_is_Invalid(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"asdasdasdasdasdasdas", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGeometryUtils_CombinePolygons_failure(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"asdasdasdasdasdasdas", "asdasdasdasdasd"}

	//When
	_, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.NotNil(t, err)
}

func TestGeometryUtils_CombinePolygons_geometry_type_MultiPolygon(t *testing.T) {
	//Given
	gm := GeometryUtils{LogLevel: logrus.DebugLevel}
	areas := []string{"frankfurt", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
