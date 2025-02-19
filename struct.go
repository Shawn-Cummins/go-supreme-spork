package main

type AutoGenerated struct {
	Introduction struct {
		Background struct {
			Text string `json:"text"`
		} `json:"Background"`
	} `json:"Introduction"`
	Geography struct {
		Location struct {
			Text string `json:"text"`
		} `json:"Location"`
		GeographicCoordinates struct {
			Text string `json:"text"`
		} `json:"Geographic coordinates"`
		MapReferences struct {
			Text string `json:"text"`
		} `json:"Map references"`
		Area struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total "`
			Land struct {
				Text string `json:"text"`
			} `json:"land"`
			Water struct {
				Text string `json:"text"`
			} `json:"water"`
			Note string `json:"note"`
		} `json:"Area"`
		AreaComparative struct {
			Text string `json:"text"`
		} `json:"Area - comparative"`
		LandBoundaries struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			BorderCountries struct {
				Text string `json:"text"`
			} `json:"border countries"`
			Note string `json:"note"`
		} `json:"Land boundaries"`
		Coastline struct {
			Text string `json:"text"`
		} `json:"Coastline"`
		MaritimeClaims struct {
			TerritorialSea struct {
				Text string `json:"text"`
			} `json:"territorial sea"`
			ContiguousZone struct {
				Text string `json:"text"`
			} `json:"contiguous zone"`
			ExclusiveEconomicZone struct {
				Text string `json:"text"`
			} `json:"exclusive economic zone"`
			ContinentalShelf struct {
				Text string `json:"text"`
			} `json:"continental shelf"`
		} `json:"Maritime claims"`
		Climate struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Climate"`
		Terrain struct {
			Text string `json:"text"`
		} `json:"Terrain"`
		Elevation struct {
			HighestPoint struct {
				Text string `json:"text"`
			} `json:"highest point"`
			LowestPoint struct {
				Text string `json:"text"`
			} `json:"lowest point"`
			MeanElevation struct {
				Text string `json:"text"`
			} `json:"mean elevation"`
			Note string `json:"note"`
		} `json:"Elevation"`
		NaturalResources struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Natural resources"`
		LandUse struct {
			AgriculturalLand struct {
				Text string `json:"text"`
			} `json:"agricultural land"`
			AgriculturalLandArableLand struct {
				Text string `json:"text"`
			} `json:"agricultural land: arable land"`
			AgriculturalLandPermanentCrops struct {
				Text string `json:"text"`
			} `json:"agricultural land: permanent crops"`
			AgriculturalLandPermanentPasture struct {
				Text string `json:"text"`
			} `json:"agricultural land: permanent pasture"`
			Forest struct {
				Text string `json:"text"`
			} `json:"forest"`
			Other struct {
				Text string `json:"text"`
			} `json:"other"`
		} `json:"Land use"`
		IrrigatedLand struct {
			Text string `json:"text"`
		} `json:"Irrigated land"`
		MajorLakesAreaSqKm struct {
			FreshWaterLakeS struct {
				Text string `json:"text"`
			} `json:"fresh water lake(s)"`
			SaltWaterLakeS struct {
				Text string `json:"text"`
			} `json:"salt water lake(s)"`
		} `json:"Major lakes (area sq km)"`
		MajorRiversByLengthInKm struct {
			Text string `json:"text"`
		} `json:"Major rivers (by length in km)"`
		MajorWatershedsAreaSqKm struct {
			Text string `json:"text"`
		} `json:"Major watersheds (area sq km)"`
		MajorAquifers struct {
			Text string `json:"text"`
		} `json:"Major aquifers"`
		PopulationDistribution struct {
			Text string `json:"text"`
		} `json:"Population distribution"`
		NaturalHazards struct {
			Text string `json:"text"`
		} `json:"Natural hazards"`
		GeographyNote struct {
			Text string `json:"text"`
		} `json:"Geography - note"`
	} `json:"Geography"`
	PeopleAndSociety struct {
		Population struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Population"`
		Nationality struct {
			Noun struct {
				Text string `json:"text"`
			} `json:"noun"`
			Adjective struct {
				Text string `json:"text"`
			} `json:"adjective"`
		} `json:"Nationality"`
		EthnicGroups struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Ethnic groups"`
		Languages struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Languages"`
		Religions struct {
			Text string `json:"text"`
		} `json:"Religions"`
		AgeStructure struct {
			Zero14Years struct {
				Text string `json:"text"`
			} `json:"0-14 years"`
			One564Years struct {
				Text string `json:"text"`
			} `json:"15-64 years"`
			Six5YearsAndOver struct {
				Text string `json:"text"`
			} `json:"65 years and over"`
		} `json:"Age structure"`
		DependencyRatios struct {
			TotalDependencyRatio struct {
				Text string `json:"text"`
			} `json:"total dependency ratio"`
			YouthDependencyRatio struct {
				Text string `json:"text"`
			} `json:"youth dependency ratio"`
			ElderlyDependencyRatio struct {
				Text string `json:"text"`
			} `json:"elderly dependency ratio"`
			PotentialSupportRatio struct {
				Text string `json:"text"`
			} `json:"potential support ratio"`
		} `json:"Dependency ratios"`
		MedianAge struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Median age"`
		PopulationGrowthRate struct {
			Text string `json:"text"`
		} `json:"Population growth rate"`
		BirthRate struct {
			Text string `json:"text"`
		} `json:"Birth rate"`
		DeathRate struct {
			Text string `json:"text"`
		} `json:"Death rate"`
		NetMigrationRate struct {
			Text string `json:"text"`
		} `json:"Net migration rate"`
		PopulationDistribution struct {
			Text string `json:"text"`
		} `json:"Population distribution"`
		Urbanization struct {
			UrbanPopulation struct {
				Text string `json:"text"`
			} `json:"urban population"`
			RateOfUrbanization struct {
				Text string `json:"text"`
			} `json:"rate of urbanization"`
		} `json:"Urbanization"`
		MajorUrbanAreasPopulation struct {
			Text string `json:"text"`
		} `json:"Major urban areas - population"`
		SexRatio struct {
			AtBirth struct {
				Text string `json:"text"`
			} `json:"at birth"`
			Zero14Years struct {
				Text string `json:"text"`
			} `json:"0-14 years"`
			One564Years struct {
				Text string `json:"text"`
			} `json:"15-64 years"`
			Six5YearsAndOver struct {
				Text string `json:"text"`
			} `json:"65 years and over"`
			TotalPopulation struct {
				Text string `json:"text"`
			} `json:"total population"`
		} `json:"Sex ratio"`
		MotherSMeanAgeAtFirstBirth struct {
			Text string `json:"text"`
		} `json:"Mother's mean age at first birth"`
		MaternalMortalityRatio struct {
			Text string `json:"text"`
		} `json:"Maternal mortality ratio"`
		InfantMortalityRate struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Infant mortality rate"`
		LifeExpectancyAtBirth struct {
			TotalPopulation struct {
				Text string `json:"text"`
			} `json:"total population"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Life expectancy at birth"`
		TotalFertilityRate struct {
			Text string `json:"text"`
		} `json:"Total fertility rate"`
		GrossReproductionRate struct {
			Text string `json:"text"`
		} `json:"Gross reproduction rate"`
		ContraceptivePrevalenceRate struct {
			Text string `json:"text"`
		} `json:"Contraceptive prevalence rate"`
		DrinkingWaterSource struct {
			ImprovedUrban struct {
				Text string `json:"text"`
			} `json:"improved: urban"`
			ImprovedRural struct {
				Text string `json:"text"`
			} `json:"improved: rural"`
			ImprovedTotal struct {
				Text string `json:"text"`
			} `json:"improved: total"`
			UnimprovedUrban struct {
				Text string `json:"text"`
			} `json:"unimproved: urban"`
			UnimprovedRural struct {
				Text string `json:"text"`
			} `json:"unimproved: rural"`
			UnimprovedTotal struct {
				Text string `json:"text"`
			} `json:"unimproved: total"`
		} `json:"Drinking water source"`
		CurrentHealthExpenditure struct {
			Text string `json:"text"`
		} `json:"Current health expenditure"`
		PhysicianDensity struct {
			Text string `json:"text"`
		} `json:"Physician density"`
		HospitalBedDensity struct {
			Text string `json:"text"`
		} `json:"Hospital bed density"`
		SanitationFacilityAccess struct {
			ImprovedUrban struct {
				Text string `json:"text"`
			} `json:"improved: urban"`
			ImprovedRural struct {
				Text string `json:"text"`
			} `json:"improved: rural"`
			ImprovedTotal struct {
				Text string `json:"text"`
			} `json:"improved: total"`
			UnimprovedUrban struct {
				Text string `json:"text"`
			} `json:"unimproved: urban"`
			UnimprovedRural struct {
				Text string `json:"text"`
			} `json:"unimproved: rural"`
			UnimprovedTotal struct {
				Text string `json:"text"`
			} `json:"unimproved: total"`
		} `json:"Sanitation facility access"`
		ObesityAdultPrevalenceRate struct {
			Text string `json:"text"`
		} `json:"Obesity - adult prevalence rate"`
		AlcoholConsumptionPerCapita struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Beer struct {
				Text string `json:"text"`
			} `json:"beer"`
			Wine struct {
				Text string `json:"text"`
			} `json:"wine"`
			Spirits struct {
				Text string `json:"text"`
			} `json:"spirits"`
			OtherAlcohols struct {
				Text string `json:"text"`
			} `json:"other alcohols"`
		} `json:"Alcohol consumption per capita"`
		TobaccoUse struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Tobacco use"`
		ChildrenUnderTheAgeOf5YearsUnderweight struct {
			Text string `json:"text"`
		} `json:"Children under the age of 5 years underweight"`
		CurrentlyMarriedWomenAges1549 struct {
			Text string `json:"text"`
		} `json:"Currently married women (ages 15-49)"`
		EducationExpenditures struct {
			Text string `json:"text"`
		} `json:"Education expenditures"`
		Literacy struct {
			TotalPopulation struct {
				Text string `json:"text"`
			} `json:"total population"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"Literacy"`
		SchoolLifeExpectancyPrimaryToTertiaryEducation struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
		} `json:"School life expectancy (primary to tertiary education)"`
	} `json:"People and Society"`
	Environment struct {
		EnvironmentCurrentIssues struct {
			Text string `json:"text"`
		} `json:"Environment - current issues"`
		EnvironmentInternationalAgreements struct {
			PartyTo struct {
				Text string `json:"text"`
			} `json:"party to"`
			SignedButNotRatified struct {
				Text string `json:"text"`
			} `json:"signed, but not ratified"`
		} `json:"Environment - international agreements"`
		Climate struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Climate"`
		LandUse struct {
			AgriculturalLand struct {
				Text string `json:"text"`
			} `json:"agricultural land"`
			AgriculturalLandArableLand struct {
				Text string `json:"text"`
			} `json:"agricultural land: arable land"`
			AgriculturalLandPermanentCrops struct {
				Text string `json:"text"`
			} `json:"agricultural land: permanent crops"`
			AgriculturalLandPermanentPasture struct {
				Text string `json:"text"`
			} `json:"agricultural land: permanent pasture"`
			Forest struct {
				Text string `json:"text"`
			} `json:"forest"`
			Other struct {
				Text string `json:"text"`
			} `json:"other"`
		} `json:"Land use"`
		Urbanization struct {
			UrbanPopulation struct {
				Text string `json:"text"`
			} `json:"urban population"`
			RateOfUrbanization struct {
				Text string `json:"text"`
			} `json:"rate of urbanization"`
		} `json:"Urbanization"`
		RevenueFromForestResources struct {
			Text string `json:"text"`
		} `json:"Revenue from forest resources"`
		RevenueFromCoal struct {
			Text string `json:"text"`
		} `json:"Revenue from coal"`
		AirPollutants struct {
			ParticulateMatterEmissions struct {
				Text string `json:"text"`
			} `json:"particulate matter emissions"`
			CarbonDioxideEmissions struct {
				Text string `json:"text"`
			} `json:"carbon dioxide emissions"`
			MethaneEmissions struct {
				Text string `json:"text"`
			} `json:"methane emissions"`
		} `json:"Air pollutants"`
		WasteAndRecycling struct {
			MunicipalSolidWasteGeneratedAnnually struct {
				Text string `json:"text"`
			} `json:"municipal solid waste generated annually"`
			MunicipalSolidWasteRecycledAnnually struct {
				Text string `json:"text"`
			} `json:"municipal solid waste recycled annually"`
			PercentOfMunicipalSolidWasteRecycled struct {
				Text string `json:"text"`
			} `json:"percent of municipal solid waste recycled"`
		} `json:"Waste and recycling"`
		MajorLakesAreaSqKm struct {
			FreshWaterLakeS struct {
				Text string `json:"text"`
			} `json:"fresh water lake(s)"`
			SaltWaterLakeS struct {
				Text string `json:"text"`
			} `json:"salt water lake(s)"`
		} `json:"Major lakes (area sq km)"`
		MajorRiversByLengthInKm struct {
			Text string `json:"text"`
		} `json:"Major rivers (by length in km)"`
		MajorWatershedsAreaSqKm struct {
			Text string `json:"text"`
		} `json:"Major watersheds (area sq km)"`
		MajorAquifers struct {
			Text string `json:"text"`
		} `json:"Major aquifers"`
		TotalWaterWithdrawal struct {
			Municipal struct {
				Text string `json:"text"`
			} `json:"municipal"`
			Industrial struct {
				Text string `json:"text"`
			} `json:"industrial"`
			Agricultural struct {
				Text string `json:"text"`
			} `json:"agricultural"`
		} `json:"Total water withdrawal"`
		TotalRenewableWaterResources struct {
			Text string `json:"text"`
		} `json:"Total renewable water resources"`
	} `json:"Environment"`
	Government struct {
		CountryName struct {
			ConventionalLongForm struct {
				Text string `json:"text"`
			} `json:"conventional long form"`
			ConventionalShortForm struct {
				Text string `json:"text"`
			} `json:"conventional short form"`
			Abbreviation struct {
				Text string `json:"text"`
			} `json:"abbreviation"`
			Etymology struct {
				Text string `json:"text"`
			} `json:"etymology"`
		} `json:"Country name"`
		GovernmentType struct {
			Text string `json:"text"`
		} `json:"Government type"`
		Capital struct {
			Name struct {
				Text string `json:"text"`
			} `json:"name"`
			GeographicCoordinates struct {
				Text string `json:"text"`
			} `json:"geographic coordinates"`
			TimeDifference struct {
				Text string `json:"text"`
			} `json:"time difference"`
			DaylightSavingTime struct {
				Text string `json:"text"`
			} `json:"daylight saving time"`
			TimeZoneNote struct {
				Text string `json:"text"`
			} `json:"time zone note"`
			Etymology struct {
				Text string `json:"text"`
			} `json:"etymology"`
		} `json:"Capital"`
		AdministrativeDivisions struct {
			Text string `json:"text"`
		} `json:"Administrative divisions"`
		DependentAreas struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Dependent areas"`
		Independence struct {
			Text string `json:"text"`
		} `json:"Independence"`
		NationalHoliday struct {
			Text string `json:"text"`
		} `json:"National holiday"`
		LegalSystem struct {
			Text string `json:"text"`
		} `json:"Legal system"`
		Constitution struct {
			History struct {
				Text string `json:"text"`
			} `json:"history"`
			Amendments struct {
				Text string `json:"text"`
			} `json:"amendments"`
		} `json:"Constitution"`
		InternationalLawOrganizationParticipation struct {
			Text string `json:"text"`
		} `json:"International law organization participation"`
		Citizenship struct {
			CitizenshipByBirth struct {
				Text string `json:"text"`
			} `json:"citizenship by birth"`
			CitizenshipByDescentOnly struct {
				Text string `json:"text"`
			} `json:"citizenship by descent only"`
			DualCitizenshipRecognized struct {
				Text string `json:"text"`
			} `json:"dual citizenship recognized"`
			ResidencyRequirementForNaturalization struct {
				Text string `json:"text"`
			} `json:"residency requirement for naturalization"`
		} `json:"Citizenship"`
		Suffrage struct {
			Text string `json:"text"`
		} `json:"Suffrage"`
		ExecutiveBranch struct {
			ChiefOfState struct {
				Text string `json:"text"`
			} `json:"chief of state"`
			HeadOfGovernment struct {
				Text string `json:"text"`
			} `json:"head of government"`
			Cabinet struct {
				Text string `json:"text"`
			} `json:"cabinet"`
			ElectionsAppointments struct {
				Text string `json:"text"`
			} `json:"elections/appointments"`
			ElectionResults struct {
				Text string `json:"text"`
			} `json:"election results"`
			Note string `json:"note"`
		} `json:"Executive branch"`
		LegislativeBranch struct {
			LegislatureName struct {
				Text string `json:"text"`
			} `json:"legislature name"`
			LegislativeStructure struct {
				Text string `json:"text"`
			} `json:"legislative structure"`
			Note string `json:"note"`
		} `json:"Legislative branch"`
		LegislativeBranchLowerChamber struct {
			ChamberName struct {
				Text string `json:"text"`
			} `json:"chamber name"`
			NumberOfSeats struct {
				Text string `json:"text"`
			} `json:"number of seats"`
			ElectoralSystem struct {
				Text string `json:"text"`
			} `json:"electoral system"`
			ScopeOfElections struct {
				Text string `json:"text"`
			} `json:"scope of elections"`
			TermInOffice struct {
				Text string `json:"text"`
			} `json:"term in office"`
			MostRecentElectionDate struct {
				Text string `json:"text"`
			} `json:"most recent election date"`
			PartiesElectedAndSeatsPerParty struct {
				Text string `json:"text"`
			} `json:"parties elected and seats per party"`
			PercentageOfWomenInChamber struct {
				Text string `json:"text"`
			} `json:"percentage of women in chamber"`
			ExpectedDateOfNextElection struct {
				Text string `json:"text"`
			} `json:"expected date of next election"`
		} `json:"Legislative branch - lower chamber"`
		LegislativeBranchUpperChamber struct {
			ChamberName struct {
				Text string `json:"text"`
			} `json:"chamber name"`
			NumberOfSeats struct {
				Text string `json:"text"`
			} `json:"number of seats"`
			ElectoralSystem struct {
				Text string `json:"text"`
			} `json:"electoral system"`
			ScopeOfElections struct {
				Text string `json:"text"`
			} `json:"scope of elections"`
			TermInOffice struct {
				Text string `json:"text"`
			} `json:"term in office"`
			MostRecentElectionDate struct {
				Text string `json:"text"`
			} `json:"most recent election date"`
			PartiesElectedAndSeatsPerParty struct {
				Text string `json:"text"`
			} `json:"parties elected and seats per party"`
			PercentageOfWomenInChamber struct {
				Text string `json:"text"`
			} `json:"percentage of women in chamber"`
			ExpectedDateOfNextElection struct {
				Text string `json:"text"`
			} `json:"expected date of next election"`
		} `json:"Legislative branch - upper chamber"`
		JudicialBranch struct {
			HighestCourtS struct {
				Text string `json:"text"`
			} `json:"highest court(s)"`
			JudgeSelectionAndTermOfOffice struct {
				Text string `json:"text"`
			} `json:"judge selection and term of office"`
			SubordinateCourts struct {
				Text string `json:"text"`
			} `json:"subordinate courts"`
			Note string `json:"note"`
		} `json:"Judicial branch"`
		PoliticalParties struct {
			Text string `json:"text"`
		} `json:"Political parties"`
		InternationalOrganizationParticipation struct {
			Text string `json:"text"`
		} `json:"International organization participation"`
		FlagDescription struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Flag description"`
		NationalSymbolS struct {
			Text string `json:"text"`
		} `json:"National symbol(s)"`
		NationalAnthem struct {
			Name struct {
				Text string `json:"text"`
			} `json:"name"`
			LyricsMusic struct {
				Text string `json:"text"`
			} `json:"lyrics/music"`
			Note string `json:"note"`
		} `json:"National anthem"`
		NationalHeritage struct {
			TotalWorldHeritageSites struct {
				Text string `json:"text"`
			} `json:"total World Heritage Sites"`
			SelectedWorldHeritageSiteLocales struct {
				Text string `json:"text"`
			} `json:"selected World Heritage Site locales"`
		} `json:"National heritage"`
	} `json:"Government"`
	Economy struct {
		EconomicOverview struct {
			Text string `json:"text"`
		} `json:"Economic overview"`
		RealGDPPurchasingPowerParity struct {
			RealGDPPurchasingPowerParity2023 struct {
				Text string `json:"text"`
			} `json:"Real GDP (purchasing power parity) 2023"`
			RealGDPPurchasingPowerParity2022 struct {
				Text string `json:"text"`
			} `json:"Real GDP (purchasing power parity) 2022"`
			RealGDPPurchasingPowerParity2021 struct {
				Text string `json:"text"`
			} `json:"Real GDP (purchasing power parity) 2021"`
			Note string `json:"note"`
		} `json:"Real GDP (purchasing power parity)"`
		RealGDPGrowthRate struct {
			RealGDPGrowthRate2023 struct {
				Text string `json:"text"`
			} `json:"Real GDP growth rate 2023"`
			RealGDPGrowthRate2022 struct {
				Text string `json:"text"`
			} `json:"Real GDP growth rate 2022"`
			RealGDPGrowthRate2021 struct {
				Text string `json:"text"`
			} `json:"Real GDP growth rate 2021"`
			Note string `json:"note"`
		} `json:"Real GDP growth rate"`
		RealGDPPerCapita struct {
			RealGDPPerCapita2023 struct {
				Text string `json:"text"`
			} `json:"Real GDP per capita 2023"`
			RealGDPPerCapita2022 struct {
				Text string `json:"text"`
			} `json:"Real GDP per capita 2022"`
			RealGDPPerCapita2021 struct {
				Text string `json:"text"`
			} `json:"Real GDP per capita 2021"`
			Note string `json:"note"`
		} `json:"Real GDP per capita"`
		GDPOfficialExchangeRate struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"GDP (official exchange rate)"`
		InflationRateConsumerPrices struct {
			InflationRateConsumerPrices2023 struct {
				Text string `json:"text"`
			} `json:"Inflation rate (consumer prices) 2023"`
			InflationRateConsumerPrices2022 struct {
				Text string `json:"text"`
			} `json:"Inflation rate (consumer prices) 2022"`
			InflationRateConsumerPrices2021 struct {
				Text string `json:"text"`
			} `json:"Inflation rate (consumer prices) 2021"`
			Note string `json:"note"`
		} `json:"Inflation rate (consumer prices)"`
		CreditRatings struct {
			FitchRating struct {
				Text string `json:"text"`
			} `json:"Fitch rating"`
			MoodySRating struct {
				Text string `json:"text"`
			} `json:"Moody's rating"`
			StandardPoorsRating struct {
				Text string `json:"text"`
			} `json:"Standard & Poors rating"`
			Note string `json:"note"`
		} `json:"Credit ratings"`
		GDPCompositionBySectorOfOrigin struct {
			Agriculture struct {
				Text string `json:"text"`
			} `json:"agriculture"`
			Industry struct {
				Text string `json:"text"`
			} `json:"industry"`
			Services struct {
				Text string `json:"text"`
			} `json:"services"`
			Note string `json:"note"`
		} `json:"GDP - composition, by sector of origin"`
		GDPCompositionByEndUse struct {
			HouseholdConsumption struct {
				Text string `json:"text"`
			} `json:"household consumption"`
			GovernmentConsumption struct {
				Text string `json:"text"`
			} `json:"government consumption"`
			InvestmentInFixedCapital struct {
				Text string `json:"text"`
			} `json:"investment in fixed capital"`
			InvestmentInInventories struct {
				Text string `json:"text"`
			} `json:"investment in inventories"`
			ExportsOfGoodsAndServices struct {
				Text string `json:"text"`
			} `json:"exports of goods and services"`
			ImportsOfGoodsAndServices struct {
				Text string `json:"text"`
			} `json:"imports of goods and services"`
			Note string `json:"note"`
		} `json:"GDP - composition, by end use"`
		AgriculturalProducts struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Agricultural products"`
		Industries struct {
			Text string `json:"text"`
		} `json:"Industries"`
		IndustrialProductionGrowthRate struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Industrial production growth rate"`
		LaborForce struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Labor force"`
		UnemploymentRate struct {
			UnemploymentRate2023 struct {
				Text string `json:"text"`
			} `json:"Unemployment rate 2023"`
			UnemploymentRate2022 struct {
				Text string `json:"text"`
			} `json:"Unemployment rate 2022"`
			UnemploymentRate2021 struct {
				Text string `json:"text"`
			} `json:"Unemployment rate 2021"`
			Note string `json:"note"`
		} `json:"Unemployment rate"`
		YouthUnemploymentRateAges1524 struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Male struct {
				Text string `json:"text"`
			} `json:"male"`
			Female struct {
				Text string `json:"text"`
			} `json:"female"`
			Note string `json:"note"`
		} `json:"Youth unemployment rate (ages 15-24)"`
		GiniIndexCoefficientDistributionOfFamilyIncome struct {
			GiniIndexCoefficientDistributionOfFamilyIncome2022 struct {
				Text string `json:"text"`
			} `json:"Gini Index coefficient - distribution of family income 2022"`
			Note string `json:"note"`
		} `json:"Gini Index coefficient - distribution of family income"`
		AverageHouseholdExpenditures struct {
			OnFood struct {
				Text string `json:"text"`
			} `json:"on food"`
			OnAlcoholAndTobacco struct {
				Text string `json:"text"`
			} `json:"on alcohol and tobacco"`
		} `json:"Average household expenditures"`
		HouseholdIncomeOrConsumptionByPercentageShare struct {
			Lowest10 struct {
				Text string `json:"text"`
			} `json:"lowest 10%"`
			Highest10 struct {
				Text string `json:"text"`
			} `json:"highest 10%"`
			Note string `json:"note"`
		} `json:"Household income or consumption by percentage share"`
		Remittances struct {
			Remittances2023 struct {
				Text string `json:"text"`
			} `json:"Remittances 2023"`
			Remittances2022 struct {
				Text string `json:"text"`
			} `json:"Remittances 2022"`
			Remittances2021 struct {
				Text string `json:"text"`
			} `json:"Remittances 2021"`
			Note string `json:"note"`
		} `json:"Remittances"`
		Budget struct {
			Revenues struct {
				Text string `json:"text"`
			} `json:"revenues"`
			Expenditures struct {
				Text string `json:"text"`
			} `json:"expenditures"`
			Note string `json:"note"`
		} `json:"Budget"`
		PublicDebt struct {
			PublicDebt2023 struct {
				Text string `json:"text"`
			} `json:"Public debt 2023"`
			Note string `json:"note"`
		} `json:"Public debt"`
		TaxesAndOtherRevenues struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Taxes and other revenues"`
		CurrentAccountBalance struct {
			CurrentAccountBalance2023 struct {
				Text string `json:"text"`
			} `json:"Current account balance 2023"`
			CurrentAccountBalance2022 struct {
				Text string `json:"text"`
			} `json:"Current account balance 2022"`
			CurrentAccountBalance2021 struct {
				Text string `json:"text"`
			} `json:"Current account balance 2021"`
			Note string `json:"note"`
		} `json:"Current account balance"`
		Exports struct {
			Exports2023 struct {
				Text string `json:"text"`
			} `json:"Exports 2023"`
			Exports2022 struct {
				Text string `json:"text"`
			} `json:"Exports 2022"`
			Exports2021 struct {
				Text string `json:"text"`
			} `json:"Exports 2021"`
			Note string `json:"note"`
		} `json:"Exports"`
		ExportsPartners struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Exports - partners"`
		ExportsCommodities struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Exports - commodities"`
		Imports struct {
			Imports2023 struct {
				Text string `json:"text"`
			} `json:"Imports 2023"`
			Imports2022 struct {
				Text string `json:"text"`
			} `json:"Imports 2022"`
			Imports2021 struct {
				Text string `json:"text"`
			} `json:"Imports 2021"`
			Note string `json:"note"`
		} `json:"Imports"`
		ImportsPartners struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Imports - partners"`
		ImportsCommodities struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Imports - commodities"`
		ReservesOfForeignExchangeAndGold struct {
			ReservesOfForeignExchangeAndGold2023 struct {
				Text string `json:"text"`
			} `json:"Reserves of foreign exchange and gold 2023"`
			ReservesOfForeignExchangeAndGold2022 struct {
				Text string `json:"text"`
			} `json:"Reserves of foreign exchange and gold 2022"`
			ReservesOfForeignExchangeAndGold2021 struct {
				Text string `json:"text"`
			} `json:"Reserves of foreign exchange and gold 2021"`
			Note string `json:"note"`
		} `json:"Reserves of foreign exchange and gold"`
		ExchangeRates struct {
			Text string `json:"text"`
		} `json:"Exchange rates"`
	} `json:"Economy"`
	Energy struct {
		ElectricityAccess struct {
			ElectrificationTotalPopulation struct {
				Text string `json:"text"`
			} `json:"electrification - total population"`
		} `json:"Electricity access"`
		Electricity struct {
			InstalledGeneratingCapacity struct {
				Text string `json:"text"`
			} `json:"installed generating capacity"`
			Consumption struct {
				Text string `json:"text"`
			} `json:"consumption"`
			Exports struct {
				Text string `json:"text"`
			} `json:"exports"`
			Imports struct {
				Text string `json:"text"`
			} `json:"imports"`
			TransmissionDistributionLosses struct {
				Text string `json:"text"`
			} `json:"transmission/distribution losses"`
		} `json:"Electricity"`
		ElectricityGenerationSources struct {
			FossilFuels struct {
				Text string `json:"text"`
			} `json:"fossil fuels"`
			Nuclear struct {
				Text string `json:"text"`
			} `json:"nuclear"`
			Solar struct {
				Text string `json:"text"`
			} `json:"solar"`
			Wind struct {
				Text string `json:"text"`
			} `json:"wind"`
			Hydroelectricity struct {
				Text string `json:"text"`
			} `json:"hydroelectricity"`
			Geothermal struct {
				Text string `json:"text"`
			} `json:"geothermal"`
			BiomassAndWaste struct {
				Text string `json:"text"`
			} `json:"biomass and waste"`
		} `json:"Electricity generation sources"`
		NuclearEnergy struct {
			NumberOfOperationalNuclearReactors struct {
				Text string `json:"text"`
			} `json:"Number of operational nuclear reactors"`
			NetCapacityOfOperationalNuclearReactors struct {
				Text string `json:"text"`
			} `json:"Net capacity of operational nuclear reactors"`
			PercentOfTotalElectricityProduction struct {
				Text string `json:"text"`
			} `json:"Percent of total electricity production"`
			NumberOfNuclearReactorsPermanentlyShutDown struct {
				Text string `json:"text"`
			} `json:"Number of nuclear reactors permanently shut down"`
		} `json:"Nuclear energy"`
		Coal struct {
			Production struct {
				Text string `json:"text"`
			} `json:"production"`
			Consumption struct {
				Text string `json:"text"`
			} `json:"consumption"`
			Exports struct {
				Text string `json:"text"`
			} `json:"exports"`
			Imports struct {
				Text string `json:"text"`
			} `json:"imports"`
			ProvenReserves struct {
				Text string `json:"text"`
			} `json:"proven reserves"`
		} `json:"Coal"`
		Petroleum struct {
			TotalPetroleumProduction struct {
				Text string `json:"text"`
			} `json:"total petroleum production"`
			RefinedPetroleumConsumption struct {
				Text string `json:"text"`
			} `json:"refined petroleum consumption"`
			CrudeOilEstimatedReserves struct {
				Text string `json:"text"`
			} `json:"crude oil estimated reserves"`
		} `json:"Petroleum"`
		NaturalGas struct {
			Production struct {
				Text string `json:"text"`
			} `json:"production"`
			Consumption struct {
				Text string `json:"text"`
			} `json:"consumption"`
			Exports struct {
				Text string `json:"text"`
			} `json:"exports"`
			Imports struct {
				Text string `json:"text"`
			} `json:"imports"`
			ProvenReserves struct {
				Text string `json:"text"`
			} `json:"proven reserves"`
		} `json:"Natural gas"`
		CarbonDioxideEmissions struct {
			TotalEmissions struct {
				Text string `json:"text"`
			} `json:"total emissions"`
			FromCoalAndMetallurgicalCoke struct {
				Text string `json:"text"`
			} `json:"from coal and metallurgical coke"`
			FromPetroleumAndOtherLiquids struct {
				Text string `json:"text"`
			} `json:"from petroleum and other liquids"`
			FromConsumedNaturalGas struct {
				Text string `json:"text"`
			} `json:"from consumed natural gas"`
		} `json:"Carbon dioxide emissions"`
		EnergyConsumptionPerCapita struct {
			TotalEnergyConsumptionPerCapita2022 struct {
				Text string `json:"text"`
			} `json:"Total energy consumption per capita 2022"`
		} `json:"Energy consumption per capita"`
	} `json:"Energy"`
	Communications struct {
		TelephonesFixedLines struct {
			TotalSubscriptions struct {
				Text string `json:"text"`
			} `json:"total subscriptions"`
			SubscriptionsPer100Inhabitants struct {
				Text string `json:"text"`
			} `json:"subscriptions per 100 inhabitants"`
		} `json:"Telephones - fixed lines"`
		TelephonesMobileCellular struct {
			TotalSubscriptions struct {
				Text string `json:"text"`
			} `json:"total subscriptions"`
			SubscriptionsPer100Inhabitants struct {
				Text string `json:"text"`
			} `json:"subscriptions per 100 inhabitants"`
		} `json:"Telephones - mobile cellular"`
		TelecommunicationSystems struct {
			GeneralAssessment struct {
				Text string `json:"text"`
			} `json:"general assessment"`
			Domestic struct {
				Text string `json:"text"`
			} `json:"domestic"`
			International struct {
				Text string `json:"text"`
			} `json:"international"`
		} `json:"Telecommunication systems"`
		BroadcastMedia struct {
			Text string `json:"text"`
		} `json:"Broadcast media"`
		InternetCountryCode struct {
			Text string `json:"text"`
		} `json:"Internet country code"`
		InternetUsers struct {
			PercentOfPopulation struct {
				Text string `json:"text"`
			} `json:"percent of population"`
		} `json:"Internet users"`
		BroadbandFixedSubscriptions struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			SubscriptionsPer100Inhabitants struct {
				Text string `json:"text"`
			} `json:"subscriptions per 100 inhabitants"`
		} `json:"Broadband - fixed subscriptions"`
		CommunicationsNote struct {
			Text string `json:"text"`
		} `json:"Communications - note"`
	} `json:"Communications"`
	Transportation struct {
		NationalAirTransportSystem struct {
			NumberOfRegisteredAirCarriers struct {
				Text string `json:"text"`
			} `json:"number of registered air carriers"`
			InventoryOfRegisteredAircraftOperatedByAirCarriers struct {
				Text string `json:"text"`
			} `json:"inventory of registered aircraft operated by air carriers"`
			AnnualPassengerTrafficOnRegisteredAirCarriers struct {
				Text string `json:"text"`
			} `json:"annual passenger traffic on registered air carriers"`
			AnnualFreightTrafficOnRegisteredAirCarriers struct {
				Text string `json:"text"`
			} `json:"annual freight traffic on registered air carriers"`
		} `json:"National air transport system"`
		CivilAircraftRegistrationCountryCodePrefix struct {
			Text string `json:"text"`
		} `json:"Civil aircraft registration country code prefix"`
		Airports struct {
			Text string `json:"text"`
		} `json:"Airports"`
		Heliports struct {
			Text string `json:"text"`
		} `json:"Heliports"`
		Pipelines struct {
			Text string `json:"text"`
		} `json:"Pipelines"`
		Railways struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			StandardGauge struct {
				Text string `json:"text"`
			} `json:"standard gauge"`
		} `json:"Railways"`
		Roadways struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			Paved struct {
				Text string `json:"text"`
			} `json:"paved"`
			Unpaved struct {
				Text string `json:"text"`
			} `json:"unpaved"`
		} `json:"Roadways"`
		Waterways struct {
			Text string `json:"text"`
		} `json:"Waterways"`
		MerchantMarine struct {
			Total struct {
				Text string `json:"text"`
			} `json:"total"`
			ByType struct {
				Text string `json:"text"`
			} `json:"by type"`
			Note string `json:"note"`
		} `json:"Merchant marine"`
		Ports struct {
			TotalPorts struct {
				Text string `json:"text"`
			} `json:"total ports"`
			Large struct {
				Text string `json:"text"`
			} `json:"large"`
			Medium struct {
				Text string `json:"text"`
			} `json:"medium"`
			Small struct {
				Text string `json:"text"`
			} `json:"small"`
			VerySmall struct {
				Text string `json:"text"`
			} `json:"very small"`
			PortsWithOilTerminals struct {
				Text string `json:"text"`
			} `json:"ports with oil terminals"`
			KeyPorts struct {
				Text string `json:"text"`
			} `json:"key ports"`
		} `json:"Ports"`
	} `json:"Transportation"`
	MilitaryAndSecurity struct {
		MilitaryAndSecurityForces struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Military and security forces"`
		MilitaryExpenditures struct {
			MilitaryExpenditures2024 struct {
				Text string `json:"text"`
			} `json:"Military Expenditures 2024"`
			MilitaryExpenditures2023 struct {
				Text string `json:"text"`
			} `json:"Military Expenditures 2023"`
			MilitaryExpenditures2022 struct {
				Text string `json:"text"`
			} `json:"Military Expenditures 2022"`
			MilitaryExpenditures2021 struct {
				Text string `json:"text"`
			} `json:"Military Expenditures 2021"`
			MilitaryExpenditures2020 struct {
				Text string `json:"text"`
			} `json:"Military Expenditures 2020"`
		} `json:"Military expenditures"`
		MilitaryAndSecurityServicePersonnelStrengths struct {
			Text string `json:"text"`
		} `json:"Military and security service personnel strengths"`
		MilitaryEquipmentInventoriesAndAcquisitions struct {
			Text string `json:"text"`
		} `json:"Military equipment inventories and acquisitions"`
		MilitaryServiceAgeAndObligation struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Military service age and obligation"`
		MilitaryDeployments struct {
			Text string `json:"text"`
		} `json:"Military deployments"`
		MilitaryNote struct {
			Text string `json:"text"`
		} `json:"Military - note"`
	} `json:"Military and Security"`
	Space struct {
		SpaceAgencyAgencies struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Space agency/agencies"`
		SpaceLaunchSiteS struct {
			Text string `json:"text"`
		} `json:"Space launch site(s)"`
		SpaceProgramOverview struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Space program overview"`
	} `json:"Space"`
	Terrorism struct {
		TerroristGroupS struct {
			Text string `json:"text"`
			Note string `json:"note"`
		} `json:"Terrorist group(s)"`
	} `json:"Terrorism"`
	TransnationalIssues struct {
		RefugeesAndInternallyDisplacedPersons struct {
			RefugeesCountryOfOrigin struct {
				Text string `json:"text"`
			} `json:"refugees (country of origin)"`
			StatelessPersons struct {
				Text string `json:"text"`
			} `json:"stateless persons"`
		} `json:"Refugees and internally displaced persons"`
		IllicitDrugs struct {
			Text string `json:"text"`
		} `json:"Illicit drugs"`
	} `json:"Transnational Issues"`
}