// Copyright © 2023 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

type EcoScoreData struct {
	Adjustments                   Adjustments    `json:"adjustments"`
	Agribalyse                    Agribalyse     `json:"agribalyse"`
	Missing                       Missing        `json:"missing"`
	MissingAgribalyseMatchWarning int            `json:"missing_agribalyse_match_warning"`
	Scores                        map[string]int `json:"scores"`
	Status                        string         `json:"status"`
}
type AggregatedOrigins struct {
	EpiScore            string         `json:"epi_score"`
	Origin              string         `json:"origin"`
	Percent             int            `json:"percent"`
	TransportationScore map[string]int `json:"transportation_score"`
}
type OriginsOfIngredients struct {
	AggregatedOrigins       []AggregatedOrigins `json:"aggregated_origins"`
	EpiScore                interface{}         `json:"epi_score"`
	EpiValue                int                 `json:"epi_value"`
	OriginsFromOriginsField []string            `json:"origins_from_origins_field"`
	TransportationScore     int                 `json:"transportation_score"`
	TransportationScores    map[string]int      `json:"transportation_scores"`
	TransportationValue     int                 `json:"transportation_value"`
	TransportationValues    map[string]int      `json:"transportation_values"`
	Value                   int                 `json:"value"`
	Values                  map[string]int      `json:"values"`
	Warning                 string              `json:"warning"`
}
type Packagings struct {
	EcoscoreMaterialScore            interface{} `json:"ecoscore_material_score"`
	EcoscoreShapeRatio               interface{} `json:"ecoscore_shape_ratio"`
	Material                         string      `json:"material"`
	NonRecyclableAndNonBiodegradable string      `json:"non_recyclable_and_non_biodegradable"`
	Shape                            string      `json:"shape"`
	MaterialShape                    string      `json:"material_shape,omitempty"`
}
type Packaging struct {
	NonRecyclableAndNonBiodegradableMaterials int          `json:"non_recyclable_and_non_biodegradable_materials"`
	Packagings                                []Packagings `json:"packagings"`
	Score                                     interface{}  `json:"score"`
	Value                                     int          `json:"value"`
	Warning                                   string       `json:"warning"`
}
type ProductionSystem struct {
	Labels  []string `json:"labels"`
	Value   int      `json:"value"`
	Warning string   `json:"warning"`
}
type ThreatenedSpecies struct {
	Warning string `json:"warning"`
}
type Adjustments struct {
	OriginsOfIngredients OriginsOfIngredients `json:"origins_of_ingredients"`
	Packaging            Packaging            `json:"packaging"`
	ProductionSystem     ProductionSystem     `json:"production_system"`
	ThreatenedSpecies    ThreatenedSpecies    `json:"threatened_species"`
}
type Missing struct {
	Categories  int `json:"categories"`
	Ingredients int `json:"ingredients"`
	Labels      int `json:"labels"`
	Origins     int `json:"origins"`
	Packagings  int `json:"packagings"`
}

type Agribalyse struct {
	AgribalyseFoodCode string      `json:"agribalyse_food_code"`
	Co2Agriculture     interface{} `json:"co2_agriculture"`
	Co2Consumption     interface{} `json:"co2_consumption"`
	Co2Distribution    interface{} `json:"co2_distribution"`
	Co2Packaging       interface{} `json:"co2_packaging"`
	Co2Processing      interface{} `json:"co2_processing"`
	Co2Total           interface{} `json:"co2_total"`
	Co2Transportation  interface{} `json:"co2_transportation"`
	Code               string      `json:"code"`
	Dqr                string      `json:"dqr"`
	EfAgriculture      interface{} `json:"ef_agriculture"`
	EfConsumption      interface{} `json:"ef_consumption"`
	EfDistribution     interface{} `json:"ef_distribution"`
	EfPackaging        interface{} `json:"ef_packaging"`
	EfProcessing       interface{} `json:"ef_processing"`
	EfTotal            interface{} `json:"ef_total"`
	EfTransportation   interface{} `json:"ef_transportation"`
	IsBeverage         int         `json:"is_beverage"`
	NameEn             string      `json:"name_en"`
	NameFr             string      `json:"name_fr"`
	Score              int         `json:"score"`
	Version            string      `json:"version"`
}
