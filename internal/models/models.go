package models

type Payload struct {
    Introduction Introduction `json:"Introduction"`
    Geography    Geography    `json:"Geography"`
    People       People       `json:"People and Society"`
}

type Introduction struct {
    Background Background `json:"Background"`
}

type Background struct {
    Text string `json:"text"`
}

type Geography struct {
    Location            TextField `json:"Location"`
    GeographicCoords    TextField `json:"Geographic coordinates"`
    MapReferences       TextField `json:"Map references"`
    Area                Area      `json:"Area"`
    AreaComparative     TextField `json:"Area - comparative"`
    LandBoundaries      LandBoundaries `json:"Land boundaries"`
    Coastline           TextField `json:"Coastline"`
    MaritimeClaims      MaritimeClaims `json:"Maritime claims"`
    Climate             TextField `json:"Climate"`
    Terrain             TextField `json:"Terrain"`
    Elevation           Elevation `json:"Elevation"`
    NaturalResources    TextField `json:"Natural resources"`
    LandUse             LandUse `json:"Land use"`
    IrrigatedLand       TextField `json:"Irrigated land"`
    MajorLakes          MajorLakes `json:"Major lakes (area sq km)"`
    MajorRivers         TextField `json:"Major rivers (by length in km)"`
    MajorWatersheds     TextField `json:"Major watersheds (area sq km)"`
    MajorAquifers       TextField `json:"Major aquifers"`
    PopulationDistribution TextField `json:"Population distribution"`
    NaturalHazards      TextField `json:"Natural hazards"`
    GeographyNote       TextField `json:"Geography - note"`
}

type TextField struct {
    Text string `json:"text"`
}

type Area struct {
    Total TextField `json:"total"`
    Land  TextField `json:"land"`
    Water TextField `json:"water"`
    Note  string    `json:"note"`
}

type LandBoundaries struct {
    Total         TextField `json:"total"`
    BorderCountries TextField `json:"border countries"`
    Note          string    `json:"note"`
}

type MaritimeClaims struct {
    TerritorialSea TextField `json:"territorial sea"`
    ContiguousZone TextField `json:"contiguous zone"`
    ExclusiveEconomicZone TextField `json:"exclusive economic zone"`
    ContinentalShelf TextField `json:"continental shelf"`
}

type Elevation struct {
    HighestPoint TextField `json:"highest point"`
    LowestPoint  TextField `json:"lowest point"`
    MeanElevation TextField `json:"mean elevation"`
    Note         string    `json:"note"`
}

type LandUse struct {
    AgriculturalLand AgriculturalLand `json:"agricultural land"`
    Forest           TextField `json:"forest"`
    Other            TextField `json:"other"`
}

type AgriculturalLand struct {
    ArableLand TextField `json:"agricultural land: arable land"`
    PermanentCrops TextField `json:"agricultural land: permanent crops"`
    PermanentPasture TextField `json:"agricultural land: permanent pasture"`
}

type MajorLakes struct {
    FreshWater TextField `json:"fresh water lake(s)"`
    SaltWater  TextField `json:"salt water lake(s)"`
}

type People struct {
    Population Population `json:"Population"`
    Nationality Nationality `json:"Nationality"`
    EthnicGroups TextField `json:"Ethnic groups"`
    Languages TextField `json:"Languages"`
    Religions TextField `json:"Religions"`
    AgeStructure AgeStructure `json:"Age structure"`
    DependencyRatios DependencyRatios `json:"Dependency ratios"`
    MedianAge MedianAge `json:"Median age"`
    PopulationGrowthRate TextField `json:"Population growth rate"`
    BirthRate TextField `json:"Birth rate"`
    DeathRate TextField `json:"Death rate"`
    NetMigrationRate TextField `json:"Net migration rate"`
    PopulationDistribution TextField `json:"Population distribution"`
    Urbanization Urbanization `json:"Urbanization"`
    MajorUrbanAreas TextField `json:"Major urban areas - population"`
    SexRatio SexRatio `json:"Sex ratio"`
    MothersMeanAgeAtFirstBirth TextField `json:"Mother's mean age at first birth"`
    MaternalMortalityRatio TextField `json:"Maternal mortality ratio"`
    InfantMortalityRate InfantMortalityRate `json:"Infant mortality rate"`
    LifeExpectancyAtBirth LifeExpectancyAtBirth `json:"Life expectancy at birth"`
    TotalFertilityRate TextField `json:"Total fertility rate"`
    GrossReproductionRate TextField `json:"Gross reproduction rate"`
    ContraceptivePrevalenceRate TextField `json:"Contraceptive prevalence rate"`
    DrinkingWaterSource DrinkingWaterSource `json:"Drinking water source"`
    CurrentHealthExpenditure TextField `json:"Current health expenditure"`
    PhysicianDensity TextField `json:"Physician density"`
    HospitalBedDensity TextField `json:"Hospital bed density"`
    SanitationFacilityAccess SanitationFacilityAccess `json:"Sanitation facility access"`
    ObesityAdultPrevalenceRate TextField `json:"Obesity - adult prevalence rate"`
    AlcoholConsumptionPerCapita AlcoholConsumptionPerCapita `json:"Alcohol consumption per capita"`
    TobaccoUse TobaccoUse `json:"Tobacco use"`
    ChildrenUnderFiveUnderweight TextField `json:"Children under the age of 5 years underweight"`
    CurrentlyMarriedWomen TextField `json:"Currently married women (ages 15-49)"`
    EducationExpenditures TextField `json:"Education expenditures"`
    Literacy Literacy `json:"Literacy"`
    SchoolLifeExpectancy TextField `json:"School life expectancy (primary to tertiary education)"`
}

type Population struct {
    Total TextField `json:"total"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}

type Nationality struct {
    Noun TextField `json:"noun"`
    Adjective TextField `json:"adjective"`
}

type AgeStructure struct {
    Age0To14 TextField `json:"0-14 years"`
    Age15To64 TextField `json:"15-64 years"`
    Age65AndOver TextField `json:"65 years and over"`
}

type DependencyRatios struct {
    TotalDependencyRatio TextField `json:"total dependency ratio"`
    YouthDependencyRatio TextField `json:"youth dependency ratio"`
    ElderlyDependencyRatio TextField `json:"elderly dependency ratio"`
    PotentialSupportRatio TextField `json:"potential support ratio"`
}

type MedianAge struct {
    Total TextField `json:"total"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}

type Urbanization struct {
    UrbanPopulation TextField `json:"urban population"`
    RateOfUrbanization TextField `json:"rate of urbanization"`
}

type SexRatio struct {
    AtBirth TextField `json:"at birth"`
    Age0To14 TextField `json:"0-14 years"`
    Age15To64 TextField `json:"15-64 years"`
    Age65AndOver TextField `json:"65 years and over"`
    TotalPopulation TextField `json:"total population"`
}

type InfantMortalityRate struct {
    Total TextField `json:"total"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}

type LifeExpectancyAtBirth struct {
    Total TextField `json:"total population"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}

type DrinkingWaterSource struct {
    Improved Improved `json:"improved"`
    Unimproved Unimproved `json:"unimproved"`
}

type Improved struct {
    Urban TextField `json:"urban"`
    Rural TextField `json:"rural"`
    Total TextField `json:"total"`
}

type Unimproved struct {
    Urban TextField `json:"urban"`
    Rural TextField `json:"rural"`
    Total TextField `json:"total"`
}

type SanitationFacilityAccess struct {
    Improved Improved `json:"improved"`
    Unimproved Unimproved `json:"unimproved"`
}

type AlcoholConsumptionPerCapita struct {
    Total TextField `json:"total"`
    Beer  TextField `json:"beer"`
    Wine  TextField `json:"wine"`
    Spirits TextField `json:"spirits"`
    OtherAlcohols TextField `json:"other alcohols"`
}

type TobaccoUse struct {
    Total TextField `json:"total"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}

type Literacy struct {
    Total TextField `json:"total population"`
    Male  TextField `json:"male"`
    Female TextField `json:"female"`
}