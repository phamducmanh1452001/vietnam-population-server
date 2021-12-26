fetch("https://www.phorifai.xyz/api/provinces?page=1&limit=63")
.then(response=>response.json())
.then(dataPopulation=>{
    $('#populationCountry').html(dataPopulation.population);
    $('#areaCountry').html(dataPopulation.area);
    let density = dataPopulation.population/dataPopulation.area;
    $('#density').html(density.toFixed(2));
})
fetch("https://www.phorifai.xyz/api/gender-chart")
.then(response=>response.json())
.then(dataGender=>{
    let populationMale = dataGender.nam;
    let populationFemale = dataGender.nu;
    let percentMale = populationMale*100/(populationMale+populationFemale);
    $('#percentMale').html(percentMale.toFixed(2));
})

if(localStorage.getItem('level')==1){
    let provinceCode = localStorage.getItem('user');
    fetch("https://www.phorifai.xyz/api/districts?province_code="+provinceCode+"&page=1&limit=63")
    .then(response=>response.json())
    .then(dataProvince=>{
        let populationIndex = dataProvince.population;
        let areaIndex = dataProvince.area;
        let densityPopulation = populationIndex/areaIndex;
        $('#dataPopulation').html(populationIndex);
        $('#areaIndex').html(areaIndex.toFixed(2));
        $('#densityPopulation').html(densityPopulation.toFixed(2));
        $('#amountIndex').html(dataProvince.amount);
        $('#levelIndex').html("các quận huyện thuộc "+dataProvince.area_name);
    })

    fetch("https://www.phorifai.xyz/api/gender-chart?code="+provinceCode)
    .then(response=>response.json())
    .then(dataGender=>{
        let populationMale = dataGender.nam;
        let populationFemale = dataGender.nu;
        let percentMale = populationMale*100/(populationMale+populationFemale);
        $('#percentMaleIndex').html(percentMale.toFixed(2));
        
    })
}
else if(localStorage.getItem('level')==2){
    let provinceCode = localStorage.getItem('user');
    fetch("https://www.phorifai.xyz/api/wards?district_code="+provinceCode+"&page=1&limit=63")
    .then(response=>response.json())
    .then(dataProvince=>{
        let populationIndex = dataProvince.population;
        let areaIndex = dataProvince.area;
        let densityPopulation = populationIndex/areaIndex;
        $('#dataPopulation').html(populationIndex);
        $('#areaIndex').html(areaIndex.toFixed(2));
        $('#densityPopulation').html(densityPopulation.toFixed(2));
        $('#amountIndex').html(dataProvince.amount);
        $('#levelIndex').html("các quận huyện thuộc "+dataProvince.area_name);
    })

    fetch("https://www.phorifai.xyz/api/gender-chart?code="+provinceCode)
    .then(response=>response.json())
    .then(dataGender=>{
        let populationMale = dataGender.nam;
        let populationFemale = dataGender.nu;
        let percentMale = populationMale*100/(populationMale+populationFemale);
        $('#percentMaleIndex').html(percentMale.toFixed(2));
    })
}
else{
    let provinceCode = localStorage.getItem('user');
    fetch("https://www.phorifai.xyz/api/wards?district_code="+provinceCode+"&page=1&limit=63")
    .then(response=>response.json())
    .then(dataProvince=>{
        let populationIndex = dataProvince.population;
        let areaIndex = dataProvince.area;
        let densityPopulation = populationIndex/areaIndex;
        $('#dataPopulation').html(populationIndex);
        $('#areaIndex').html(areaIndex.toFixed(2));
        $('#densityPopulation').html(densityPopulation.toFixed(2));
        $('#amountIndex').html(dataProvince.amount);
        $('#levelIndex').html("các quận huyện thuộc "+dataProvince.area_name);
    })

    fetch("https://www.phorifai.xyz/api/gender-chart?code="+provinceCode)
    .then(response=>response.json())
    .then(dataGender=>{
        let populationMale = dataGender.nam;
        let populationFemale = dataGender.nu;
        let percentMale = populationMale*100/(populationMale+populationFemale);
        $('#percentMaleIndex').html(percentMale.toFixed(2));
    })   
}