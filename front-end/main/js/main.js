var api_url =  "https://www.phorifai.xyz/api/provinces?page=1&limit=63";
fetch(api_url)
.then(response=>response.json())
.then(data=>{
    // Lấy dân số của Việt Nam
    var populationCountry = data.population;
    $('#populationCountry').html(populationCountry);
    // Lấy dân số top 30 tỉnh thành nhiều nhất
    var provinces= data.data;
    provinces.sort(function(a, b){
        return b.population - a.population;
    })
    for(let r =0;r < 10;r++){
        let v = `<div class="sec-province">
                 <span class="top-rank">${r+1}</span>
                    <span class="top-province" onclick="clickProvince('${provinces[r].code}')">${provinces[r].name}</span>
                    <span class="population-province">${provinces[r].population}</span>                                
                    </div>`
        $('#province-item10').append(v);
    }
    for(let r =10;r < 20;r++){
        let v = `<div class="sec-province">
                 <span class="top-rank">${r+1}</span>
                    <span class="top-province" onclick="clickProvince('${provinces[r].code}')">${provinces[r].name}</span>
                    <span class="population-province">${provinces[r].population}</span>                                
                    </div>`
        $('#province-item20').append(v);
    }
    for(let r =20;r < 30;r++){
        let v = `<div class="sec-province">
                 <span class="top-rank">${r+1}</span>
                    <span class="top-province" onclick="clickProvince('${provinces[r].code}')">${provinces[r].name}</span>
                    <span class="population-province">${provinces[r].population}</span>                                
                    </div>`
        $('#province-item30').append(v);
    }
    for(let r = 0; r < 10; r++){
        let v = `
    <tr>
        <td>2010</td>
        <td>8045</td>
        <td>1.1%</td>
        <td>80</td>
        <td>400</td>
        <td>30.2</td>
        <td>51.2</td>
        <td>48.8</td>
        <td>32%</td>
        <td>2600</td>
        <td>15</td>
      </tr>`
    $("#population-detail").append(v);
    }
    // Hiển thị tất cả tỉnh thành sau khi ấn nút xem tất cả
    for(let r =0;r < provinces.length;r++){
        let density = provinces[r].population/provinces[r].area;
        
        let v = `
        <tr class = "province-item">
            <td>${r+1}</td>
            <td class="top-province" id = "${provinces[r].code}" onclick="clickProvince('${provinces[r].code}')">${provinces[r].name}</td>
            <td>${provinces[r].population}</td>
            <td>${provinces[r].area}</td>
            <td>${density.toFixed(2)}</td>
            <td>37</td>
            <td>48.8</td>
            <td>37.7</td>
            <td>3675</td>
            <td>15</td>
        </tr>`
        $('#provinceAll-table').append(v);
    }
})
// Click xem tất cả các tỉnh thành
$('#all-province').click(function(){
   $('#population-container').load('province.html');
});

