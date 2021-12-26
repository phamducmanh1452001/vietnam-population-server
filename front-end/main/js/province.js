// Hàm search cho tất cả ô input
function search(inputSearch, tableSearch){
    $(inputSearch).on("keyup", function(){
        var provinceSearch = $(this).val().toLowerCase();
        $(tableSearch).filter(function(){
            $(this).toggle($(this).text().toLowerCase().indexOf(provinceSearch) > -1);
        });
    });
}
// Search tỉnh thành 
search('#search-province','.province-item');
function clickProvince(provinceCode){
    fetch("https://www.phorifai.xyz/api/districts?province_code="+provinceCode+"&page=1&limit=30")
    .then(response=>response.json())
    .then(dataDistrict=>{
        var districts = dataDistrict.data;
        districts.sort(function(a, b){
            return b.population - a.population;
        })
        // tinhs tỷ lệ nam nữ cho các huyện
        for(let r =0;r < districts.length;r++){
            fetch("https://www.phorifai.xyz/api/gender-chart?code="+districts[r].code)
            .then(response=>response.json())
            .then(dataGender=>{
                var percentMale = dataGender.nam *100/districts[r].population;
                var percentFemale = 100-percentMale.toFixed(2);
                let v = `
                <tr class = "province-item district-item">
                    <td class="top-province" id = "${districts[r].code}" onclick="clickDistrict('${districts[r].code}', '${provinceCode}')">${districts[r].name}</td>
                    <td>${districts[r].population}</td>
                    <td>${districts[r].area}</td>
                    <td>${(districts[r].population/districts[r].area).toFixed(2)}</td>
                    <td>51.2</td>
                    <td id="percentMale">${percentMale.toFixed(2)}</td>
                    <td id="percentFemale">${percentFemale.toFixed(2)}</td>
                    <td>3675</td>
                    <td>15</td>
                </tr>` 
            $('#districtAll-table').append(v);
            })
        }  
    
        $('.areaName').html(dataDistrict.area_name);       
        $('.areaPopulation').html(dataDistrict.population);
        var api_url =  "https://www.phorifai.xyz/api/provinces?page=1&limit=63";
        fetch(api_url)
        .then(response=>response.json())
        .then(dataProvince=>{
            let populationCountry = dataProvince.population;
            let populationProvince = dataDistrict.population 
            let area = dataDistrict.area;
            let percentPopulationCountry = populationProvince*100 / populationCountry;
            let density = populationProvince/area;            
            $('#percentPopulation').html(percentPopulationCountry.toFixed(2));
            $('#area').html(area.toFixed(2));
            $('#density').html(density.toFixed(2));
            $('.levelName').html('huyện');
        })
    })
    // TÍnh tỷ lệ giới tính
    var chartEle = '<canvas id="myCharty" class="chart" ></canvas>';
    $('#chartAge').append(chartEle);
    fetch("https://www.phorifai.xyz/api/gender-chart?code="+provinceCode)
    .then(response=>response.json())
    .then(dataGender=>{
        let populationAll = dataGender.nam + dataGender.nu;
        let percentMale = dataGender.nam*100/populationAll;
        $('.percentMale').html(percentMale.toFixed(2));

        // hiển thị biểu đồ theo giới tính
        let xValues = ["Nam", "Nữ"];
        let yValues = [dataGender.nam, dataGender.nu];
        let barColors = [
          "#17a2b8",
          "#6f42c1"
        ];
        // Biểu đồ tỷ lệ giới tính
        new Chart("myCharty", {
          type: "pie",
          data: {
            labels: xValues,
            datasets: [{
              backgroundColor: barColors,
              data: yValues
            }]
          },
          options: {
            title: {
              display: true,
              text: "Biểu đồ dân số theo giới tính",
              position: "bottom",
              fontSize: 20
            },
            legend:{
              display: true,
              labels: {
              }
            },
            layout:{
              padding:{
                top: 20
              }      
            }
          }
        });
    })

    //Biều đồ dân số theo tôn giáo
    var chartEleRe = '<canvas id="chartreligion" class="chart" ></canvas>';
    $('#chartReligion').append(chartEleRe);
    fetch("https://www.phorifai.xyz/api/religion-chart?code="+provinceCode)
    .then(response=>response.json())
    .then(dataReligion=>{
        var xValues = ["Cao Đài", "Công Giáo", "Hòa Hảo", "Hồi Giáo", "Khác", "Không", "Phật giáo","Tin lành","Tôn giáo dân gian"];
        // Dân số qua các dân tộc
        var yValues = [dataReligion.cao_dai, dataReligion.cong_giao , dataReligion.hoa_hao, dataReligion.hoi_giao, dataReligion.khac, dataReligion.khong,dataReligion.phat_giao,dataReligion.tin_lanh,dataReligion.ton_giao_dan_gian];
        var barColors = [
            "#5d767e",
            "#00aba9",
            "#2b5797",
            "#e8c3b9",
            "#1e7145",
            "#ffc107",
            "#fd7e14",
            "#6c757d",
            "#b91d47"
        ];

        new Chart("chartreligion", {
            type: "pie",
            data: {
            labels: xValues,
            datasets: [{
                backgroundColor: barColors,
                data: yValues
            }]
            },
            options: {
            title: {
                display: true,
                text: "Biểu đồ cơ cấu theo tôn giáo của nước ta năm 2020",
                position: "bottom",
                fontSize: 20
            },
            legend:{
                display: true,
            }
            }
        });
    })

    // biểu đồ theo tuổi
    var chartEleAge = '<canvas id="chart_age"></canvas>';
    $('#chart_age').append(chartEleAge);
        fetch("https://www.phorifai.xyz/api/age-chart?code="+provinceCode)
        .then(response=>response.json())
        .then(dataAge=>{
        // Biểu đồ giới tính
            console.log(dataAge);
            var xValues = ["0-10", "10-20", "20-30", "30-40", "40-50","50-60","60-70","70-80","80-90","90 trở lên"];
            var yValues = [];
        for(let key in dataAge){
            yValues.push(dataAge[key]);
        }
        new Chart("chart_age", {
            type: "bar",
            data: {
            labels: xValues,  
            datasets: [{
                label: 'Dân số',
                backgroundColor: "#2b5797",
                data: yValues
            }]
            },
            options: {
            legend: {
                display: true,
            },
            title: {
                display: true,
                text: "Biểu đồ phân chia dân số theo độ tuổi",
                position: 'bottom',
                fontSize: 20
            }
            }
        });
    })
    // load trang cấp tỉnh
    $('#population-container').load('district.html');  
    // load button back
    let btn_back = '<i class="fas fa-arrow-left"></i>'
    $('#back').html(btn_back);
    $('#back').click(function(){
        console.log(10);
        $('.all').load('mainpage.html');
    })
}

// search huyện của mỗi tỉnh thành 
search('#search-district', '.district-item');